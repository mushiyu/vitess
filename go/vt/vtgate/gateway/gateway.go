/*
Copyright 2017 Google Inc.

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

// Package gateway contains the routing layer of vtgate. A Gateway can take
// a query targeted to a keyspace/shard/tablet_type and send it off.
package gateway

import (
	"flag"
	"fmt"
	"time"

	"github.com/mushiyu/vitess/go/vt/log"
	"golang.org/x/net/context"

	"github.com/mushiyu/vitess/go/vt/discovery"
	"github.com/mushiyu/vitess/go/vt/srvtopo"
	"github.com/mushiyu/vitess/go/vt/vttablet/queryservice"

	querypb "github.com/mushiyu/vitess/go/vt/proto/query"
	topodatapb "github.com/mushiyu/vitess/go/vt/proto/topodata"
)

// This file contains the Gateway interface definition, and the
// implementations registry.

var (
	implementation       = flag.String("gateway_implementation", "discoverygateway", "The implementation of gateway")
	initialTabletTimeout = flag.Duration("gateway_initial_tablet_timeout", 30*time.Second, "At startup, the gateway will wait up to that duration to get one tablet per keyspace/shard/tablettype")
)

// A Gateway is the query processing module for each shard,
// which is used by ScatterConn.
type Gateway interface {
	// TODO(alainjobart) The QueryService part of this interface
	// will be removed soon, in favor of the TargetStats part (that
	// returns a QueryService)
	queryservice.QueryService

	// srvtopo.TargetStats allows this Gateway to resolve a Target
	// into a QueryService. It is used by the srvtopo.Resolver object.
	srvtopo.TargetStats

	// WaitForTablets asks the gateway to wait for the provided
	// tablets types to be available. It the context is canceled
	// before the end, it should return ctx.Err().
	// The error returned will have specific effects:
	// - nil: keep going with startup.
	// - context.DeadlineExceeded: log a warning that we didn't get
	//   all tablets, and keep going with startup.
	// - any other error: log.Fatalf out.
	WaitForTablets(ctx context.Context, tabletTypesToWait []topodatapb.TabletType) error

	// RegisterStats registers exported stats for the gateway
	RegisterStats()

	// CacheStatus returns a list of TabletCacheStatus per shard / tablet type.
	CacheStatus() TabletCacheStatusList
}

// Creator is the factory method which can create the actual gateway object.
type Creator func(hc discovery.HealthCheck, serv srvtopo.Server, cell string, retryCount int) Gateway

var creators = make(map[string]Creator)

// RegisterCreator registers a Creator with given name.
func RegisterCreator(name string, gc Creator) {
	if _, ok := creators[name]; ok {
		log.Fatalf("Gateway %s already exists", name)
	}
	creators[name] = gc
}

// GetCreator returns the Creator specified by the gateway_implementation flag.
func GetCreator() Creator {
	gc, ok := creators[*implementation]
	if !ok {
		log.Exitf("No gateway registered as %s", *implementation)
	}
	return gc
}

// WaitForTablets is a helper method to wait for the provided tablets,
// up until the *initialTabletTimeout. It will log what it is doing.
// Note it has the same name as the Gateway's interface method, as it
// just calls it.
func WaitForTablets(gw Gateway, tabletTypesToWait []topodatapb.TabletType) error {
	log.Infof("Gateway waiting for serving tablets...")
	ctx, cancel := context.WithTimeout(context.Background(), *initialTabletTimeout)
	defer cancel()

	err := gw.WaitForTablets(ctx, tabletTypesToWait)
	switch err {
	case nil:
		// Log so we know everything is fine.
		log.Infof("Waiting for tablets completed")
	case context.DeadlineExceeded:
		// In this scenario, we were able to reach the
		// topology service, but some tablets may not be
		// ready. We just warn and keep going.
		log.Warningf("Timeout waiting for all keyspaces / shards to have healthy tablets, may be in degraded mode")
		err = nil
	default:
		// Nothing to do here, the caller will log.Fatalf.
	}
	return err
}

// StreamHealthFromTargetStatsListener responds to a StreamHealth
// streaming RPC using a srvtopo.TargetStatsListener implementation.
func StreamHealthFromTargetStatsListener(ctx context.Context, l srvtopo.TargetStatsListener, callback func(*querypb.StreamHealthResponse) error) error {
	// Subscribe to the TargetStatsListener aggregate stats.
	id, entries, c, err := l.Subscribe()
	if err != nil {
		return err
	}
	defer func() {
		// Unsubscribe so we don't receive more updates, and
		// drain the channel.
		l.Unsubscribe(id)
		for range c {
		}
	}()

	// Send all current entries.
	for _, e := range entries {
		shr := &querypb.StreamHealthResponse{
			Target:                              e.Target,
			TabletExternallyReparentedTimestamp: e.TabletExternallyReparentedTimestamp,
			AggregateStats:                      e.Stats,
		}
		if err := callback(shr); err != nil {
			return err
		}
	}

	// Now listen for updates, or the end of the connection.
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case e, ok := <-c:
			if !ok {
				// Channel is closed, should never happen.
				return fmt.Errorf("channel closed")
			}
			shr := &querypb.StreamHealthResponse{
				Target:                              e.Target,
				TabletExternallyReparentedTimestamp: e.TabletExternallyReparentedTimestamp,
				AggregateStats:                      e.Stats,
			}
			if err := callback(shr); err != nil {
				return err
			}
		}
	}
}
