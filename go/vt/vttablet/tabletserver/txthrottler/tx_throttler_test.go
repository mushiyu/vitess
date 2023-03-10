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

package txthrottler

// Commands to generate the mocks for this test.
//go:generate mockgen -destination mock_healthcheck_test.go -package txthrottler github.com/mushiyu/vitess/go/vt/discovery HealthCheck
//go:generate mockgen -destination mock_throttler_test.go -package txthrottler github.com/mushiyu/vitess/go/vt/vttablet/tabletserver/txthrottler ThrottlerInterface
//go:generate mockgen -destination mock_topology_watcher_test.go -package txthrottler github.com/mushiyu/vitess/go/vt/vttablet/tabletserver/txthrottler TopologyWatcherInterface

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/mushiyu/vitess/go/vt/discovery"
	"github.com/mushiyu/vitess/go/vt/topo"
	"github.com/mushiyu/vitess/go/vt/topo/memorytopo"
	"github.com/mushiyu/vitess/go/vt/vttablet/tabletserver/tabletenv"

	querypb "github.com/mushiyu/vitess/go/vt/proto/query"
	topodatapb "github.com/mushiyu/vitess/go/vt/proto/topodata"
)

func TestDisabledThrottler(t *testing.T) {
	oldConfig := tabletenv.Config
	defer func() { tabletenv.Config = oldConfig }()
	tabletenv.Config.EnableTxThrottler = false
	throttler := CreateTxThrottlerFromTabletConfig(nil)
	if err := throttler.Open("keyspace", "shard"); err != nil {
		t.Fatalf("want: nil, got: %v", err)
	}
	if result := throttler.Throttle(); result != false {
		t.Errorf("want: false, got: %v", result)
	}
	throttler.Close()
}

func TestEnabledThrottler(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	defer resetTxThrottlerFactories()
	ts := memorytopo.NewServer("cell1", "cell2")

	mockHealthCheck := NewMockHealthCheck(mockCtrl)
	var hcListener discovery.HealthCheckStatsListener
	hcCall1 := mockHealthCheck.EXPECT().SetListener(gomock.Any(), false /* sendDownEvents */)
	hcCall1.Do(func(listener discovery.HealthCheckStatsListener, sendDownEvents bool) {
		// Record the listener we're given.
		hcListener = listener
	})
	hcCall2 := mockHealthCheck.EXPECT().Close()
	hcCall2.After(hcCall1)
	healthCheckFactory = func() discovery.HealthCheck { return mockHealthCheck }

	topologyWatcherFactory = func(topoServer *topo.Server, tr discovery.TabletRecorder, cell, keyspace, shard string, refreshInterval time.Duration, topoReadConcurrency int) TopologyWatcherInterface {
		if ts != topoServer {
			t.Errorf("want: %v, got: %v", ts, topoServer)
		}
		if cell != "cell1" && cell != "cell2" {
			t.Errorf("want: cell1 or cell2, got: %v", cell)
		}
		if keyspace != "keyspace" {
			t.Errorf("want: keyspace, got: %v", keyspace)
		}
		if shard != "shard" {
			t.Errorf("want: shard, got: %v", shard)
		}
		result := NewMockTopologyWatcherInterface(mockCtrl)
		result.EXPECT().Stop()
		return result
	}

	mockThrottler := NewMockThrottlerInterface(mockCtrl)
	throttlerFactory = func(name, unit string, threadCount int, maxRate, maxReplicationLag int64) (ThrottlerInterface, error) {
		if threadCount != 1 {
			t.Errorf("want: 1, got: %v", threadCount)
		}
		return mockThrottler, nil
	}

	call0 := mockThrottler.EXPECT().UpdateConfiguration(gomock.Any(), true /* copyZeroValues */)
	call1 := mockThrottler.EXPECT().Throttle(0)
	call1.Return(0 * time.Second)
	tabletStats := &discovery.TabletStats{
		Target: &querypb.Target{
			TabletType: topodatapb.TabletType_REPLICA,
		},
	}
	call2 := mockThrottler.EXPECT().RecordReplicationLag(gomock.Any(), tabletStats)
	call3 := mockThrottler.EXPECT().Throttle(0)
	call3.Return(1 * time.Second)
	call4 := mockThrottler.EXPECT().Close()
	call1.After(call0)
	call2.After(call1)
	call3.After(call2)
	call4.After(call3)

	oldConfig := tabletenv.Config
	defer func() { tabletenv.Config = oldConfig }()
	tabletenv.Config.EnableTxThrottler = true
	tabletenv.Config.TxThrottlerHealthCheckCells = []string{"cell1", "cell2"}

	throttler, err := tryCreateTxThrottler(ts)
	if err != nil {
		t.Fatalf("want: nil, got: %v", err)
	}
	if err := throttler.Open("keyspace", "shard"); err != nil {
		t.Fatalf("want: nil, got: %v", err)
	}
	if result := throttler.Throttle(); result != false {
		t.Errorf("want: false, got: %v", result)
	}
	hcListener.StatsUpdate(tabletStats)
	rdonlyTabletStats := &discovery.TabletStats{
		Target: &querypb.Target{
			TabletType: topodatapb.TabletType_RDONLY,
		},
	}
	// This call should not be forwarded to the go/vt/throttler.Throttler object.
	hcListener.StatsUpdate(rdonlyTabletStats)
	// The second throttle call should reject.
	if result := throttler.Throttle(); result != true {
		t.Errorf("want: true, got: %v", result)
	}
	throttler.Close()
}
