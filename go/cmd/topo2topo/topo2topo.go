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

package main

import (
	"flag"

	"github.com/mushiyu/vitess/go/exit"
	"github.com/mushiyu/vitess/go/vt/log"
	"github.com/mushiyu/vitess/go/vt/logutil"
	"github.com/mushiyu/vitess/go/vt/topo"
	"github.com/mushiyu/vitess/go/vt/topo/helpers"
	"golang.org/x/net/context"
)

var (
	fromImplementation = flag.String("from_implementation", "", "topology implementation to copy data from")
	fromServerAddress  = flag.String("from_server", "", "topology server address to copy data from")
	fromRoot           = flag.String("from_root", "", "topology server root to copy data from")

	toImplementation = flag.String("to_implementation", "", "topology implementation to copy data to")
	toServerAddress  = flag.String("to_server", "", "topology server address to copy data to")
	toRoot           = flag.String("to_root", "", "topology server root to copy data to")

	doKeyspaces         = flag.Bool("do-keyspaces", false, "copies the keyspace information")
	doShards            = flag.Bool("do-shards", false, "copies the shard information")
	doShardReplications = flag.Bool("do-shard-replications", false, "copies the shard replication information")
	doTablets           = flag.Bool("do-tablets", false, "copies the tablet information")
)

func main() {
	defer exit.RecoverAll()
	defer logutil.Flush()

	flag.Parse()
	args := flag.Args()
	if len(args) != 0 {
		flag.Usage()
		log.Exitf("topo2topo doesn't take any parameter.")
	}

	fromTS, err := topo.OpenServer(*fromImplementation, *fromServerAddress, *fromRoot)
	if err != nil {
		log.Exitf("Cannot open 'from' topo %v: %v", *fromImplementation, err)
	}
	toTS, err := topo.OpenServer(*toImplementation, *toServerAddress, *toRoot)
	if err != nil {
		log.Exitf("Cannot open 'to' topo %v: %v", *toImplementation, err)
	}

	ctx := context.Background()

	if *doKeyspaces {
		helpers.CopyKeyspaces(ctx, fromTS, toTS)
	}
	if *doShards {
		helpers.CopyShards(ctx, fromTS, toTS)
	}
	if *doShardReplications {
		helpers.CopyShardReplications(ctx, fromTS, toTS)
	}
	if *doTablets {
		helpers.CopyTablets(ctx, fromTS, toTS)
	}
}
