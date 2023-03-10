/*
Copyright 2017 Google Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreedto in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package heartbeat

import (
	"fmt"
	"testing"
	"time"

	"math/rand"

	"github.com/mushiyu/vitess/go/mysql/fakesqldb"
	"github.com/mushiyu/vitess/go/sqlescape"
	"github.com/mushiyu/vitess/go/sqltypes"
	"github.com/mushiyu/vitess/go/vt/dbconfigs"
	"github.com/mushiyu/vitess/go/vt/vttablet/tabletserver/tabletenv"

	querypb "github.com/mushiyu/vitess/go/vt/proto/query"
)

// TestReaderReadHeartbeat tests that reading a heartbeat sets the appropriate
// fields on the object.
func TestReaderReadHeartbeat(t *testing.T) {
	db := fakesqldb.New(t)
	defer db.Close()
	tr := newReader(db, mockNowFunc)
	defer tr.Close()

	db.AddQuery(fmt.Sprintf("SELECT ts FROM %s.heartbeat WHERE keyspaceShard='%s'", tr.dbName, tr.keyspaceShard), &sqltypes.Result{
		Fields: []*querypb.Field{
			{Name: "ts", Type: sqltypes.Int64},
		},
		Rows: [][]sqltypes.Value{{
			sqltypes.NewInt64(now.Add(-10 * time.Second).UnixNano()),
		}},
	})

	cumulativeLagNs.Reset()
	readErrors.Reset()
	reads.Reset()

	tr.readHeartbeat()
	lag, err := tr.GetLatest()

	if err != nil {
		t.Fatalf("Should not be in error: %v", tr.lastKnownError)
	}
	if got, want := lag, 10*time.Second; got != want {
		t.Fatalf("wrong latest lag: got = %v, want = %v", tr.lastKnownLag, want)
	}
	if got, want := cumulativeLagNs.Get(), 10*time.Second.Nanoseconds(); got != want {
		t.Fatalf("wrong cumulative lag: got = %v, want = %v", got, want)
	}
	if got, want := reads.Get(), int64(1); got != want {
		t.Fatalf("wrong read count: got = %v, want = %v", got, want)
	}
	if got, want := readErrors.Get(), int64(0); got != want {
		t.Fatalf("wrong read error count: got = %v, want = %v", got, want)
	}
}

// TestReaderReadHeartbeatError tests that we properly account for errors
// encountered in the reading of heartbeat.
func TestReaderReadHeartbeatError(t *testing.T) {
	db := fakesqldb.New(t)
	defer db.Close()
	tr := newReader(db, mockNowFunc)
	defer tr.Close()

	cumulativeLagNs.Reset()
	readErrors.Reset()

	tr.readHeartbeat()
	lag, err := tr.GetLatest()

	if err == nil {
		t.Fatalf("Should be in error: %v", tr.lastKnownError)
	}
	if got, want := lag, 0*time.Second; got != want {
		t.Fatalf("wrong lastKnownLag: got = %v, want = %v", got, want)
	}
	if got, want := cumulativeLagNs.Get(), int64(0); got != want {
		t.Fatalf("wrong cumulative lag: got = %v, want = %v", got, want)
	}
	if got, want := readErrors.Get(), int64(1); got != want {
		t.Fatalf("wrong read error count: got = %v, want = %v", got, want)
	}
}

func newReader(db *fakesqldb.DB, nowFunc func() time.Time) *Reader {
	randID := rand.Int63()
	config := tabletenv.DefaultQsConfig
	config.HeartbeatEnable = true
	config.PoolNamePrefix = fmt.Sprintf("Pool-%d-", randID)
	dbc := dbconfigs.NewTestDBConfigs(*db.ConnParams(), *db.ConnParams(), "")

	tr := NewReader(&fakeMysqlChecker{}, config)
	tr.dbName = sqlescape.EscapeID(dbc.SidecarDBName.Get())
	tr.keyspaceShard = "test:0"
	tr.now = nowFunc
	tr.pool.Open(dbc.AppWithDB(), dbc.DbaWithDB(), dbc.AppDebugWithDB())

	return tr
}
