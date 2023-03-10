/*
Copyright 2018 The Vitess Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package vreplication

import (
	"flag"
	"fmt"
	"strconv"
	"time"

	"github.com/mushiyu/vitess/go/vt/vterrors"

	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"

	"github.com/mushiyu/vitess/go/sync2"
	"github.com/mushiyu/vitess/go/tb"
	"github.com/mushiyu/vitess/go/vt/binlog/binlogplayer"
	"github.com/mushiyu/vitess/go/vt/log"
	"github.com/mushiyu/vitess/go/vt/mysqlctl"
	"github.com/mushiyu/vitess/go/vt/topo"

	binlogdatapb "github.com/mushiyu/vitess/go/vt/proto/binlogdata"
)

var retryDelay = flag.Duration("vreplication_retry_delay", 5*time.Second, "delay before retrying a failed binlog connection")

// controller is created by Engine. Members are initialized upfront.
// There is no mutex within a controller becaust its members are
// either read-only or self-synchronized.
type controller struct {
	dbClientFactory func() binlogplayer.DBClient
	mysqld          mysqlctl.MysqlDaemon
	blpStats        *binlogplayer.Stats

	id           uint32
	source       binlogdatapb.BinlogSource
	stopPos      string
	tabletPicker *tabletPicker

	cancel context.CancelFunc
	done   chan struct{}

	// The following fields are updated after start. So, they need synchronization.
	sourceTablet sync2.AtomicString
}

// newController creates a new controller. Unless a stream is explicitly 'Stopped',
// this function launches a goroutine to perform continuous vreplication.
func newController(ctx context.Context, params map[string]string, dbClientFactory func() binlogplayer.DBClient, mysqld mysqlctl.MysqlDaemon, ts *topo.Server, cell, tabletTypesStr string, blpStats *binlogplayer.Stats) (*controller, error) {
	if blpStats == nil {
		blpStats = binlogplayer.NewStats()
	}
	ct := &controller{
		dbClientFactory: dbClientFactory,
		mysqld:          mysqld,
		blpStats:        blpStats,
		done:            make(chan struct{}),
	}

	// id
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		return nil, err
	}
	ct.id = uint32(id)

	// Nothing to do if replication is stopped.
	if params["state"] == binlogplayer.BlpStopped {
		ct.cancel = func() {}
		close(ct.done)
		return ct, nil
	}

	// source, stopPos
	if err := proto.UnmarshalText(params["source"], &ct.source); err != nil {
		return nil, err
	}
	ct.stopPos = params["stop_pos"]

	// tabletPicker
	if v, ok := params["cell"]; ok {
		cell = v
	}
	if v, ok := params["tablet_types"]; ok {
		tabletTypesStr = v
	}
	tp, err := newTabletPicker(ts, cell, ct.source.Keyspace, ct.source.Shard, tabletTypesStr)
	if err != nil {
		return nil, err
	}
	ct.tabletPicker = tp

	// cancel
	ctx, ct.cancel = context.WithCancel(ctx)

	go ct.run(ctx)

	return ct, nil
}

func (ct *controller) run(ctx context.Context) {
	defer func() {
		log.Infof("stream %v: stopped", ct.id)
		ct.tabletPicker.Close()
		close(ct.done)
	}()

	for {
		err := ct.runBlp(ctx)
		if err == nil {
			return
		}
		// Sometimes, canceled contexts get wrapped as errors.
		select {
		case <-ctx.Done():
			return
		default:
		}
		log.Warningf("stream %v: %v, retrying after %v", ct.id, err, *retryDelay)
		time.Sleep(*retryDelay)
	}
}

func (ct *controller) runBlp(ctx context.Context) (err error) {
	defer func() {
		ct.sourceTablet.Set("")
		if x := recover(); x != nil {
			log.Errorf("stream %v: caught panic: %v\n%s", ct.id, x, tb.Stack(4))
			err = fmt.Errorf("panic: %v", x)
		}
	}()

	select {
	case <-ctx.Done():
		return nil
	default:
	}

	// Call this for youtube-specific customization.
	// This should be done every time, in case mysql was restarted.
	if err := ct.mysqld.EnableBinlogPlayback(); err != nil {
		return err
	}

	dbClient := ct.dbClientFactory()
	if err := dbClient.Connect(); err != nil {
		return vterrors.Wrap(err, "can't connect to database")
	}
	defer dbClient.Close()

	tablet, err := ct.tabletPicker.Pick(ctx)
	if err != nil {
		return err
	}
	ct.sourceTablet.Set(tablet.Alias.String())

	if len(ct.source.Tables) > 0 {
		// Table names can have search patterns. Resolve them against the schema.
		tables, err := mysqlctl.ResolveTables(ct.mysqld, dbClient.DBName(), ct.source.Tables)
		if err != nil {
			return vterrors.Wrap(err, "failed to resolve table names")
		}

		player := binlogplayer.NewBinlogPlayerTables(dbClient, tablet, tables, ct.id, ct.blpStats)
		return player.ApplyBinlogEvents(ctx)
	}
	player := binlogplayer.NewBinlogPlayerKeyRange(dbClient, tablet, ct.source.KeyRange, ct.id, ct.blpStats)
	return player.ApplyBinlogEvents(ctx)
}

func (ct *controller) Stop() {
	ct.cancel()
	<-ct.done
}
