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

package worker

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/mushiyu/vitess/go/vt/discovery"
	"github.com/mushiyu/vitess/go/vt/topo/topoproto"

	topodatapb "github.com/mushiyu/vitess/go/vt/proto/topodata"
)

// TabletTracker tracks for each tablet alias how often it is currently in use
// for a streaming read query.
// By using this information, all streaming queries should be balanced across
// all available tablets.
type TabletTracker struct {
	// mu guards the fields below.
	mu sync.Mutex
	// usedTablets stores how often a tablet is currently in use.
	// The map key is the string of the TabletAlias.
	usedTablets map[string]int
}

// NewTabletTracker returns a new TabletTracker.
func NewTabletTracker() *TabletTracker {
	return &TabletTracker{
		usedTablets: make(map[string]int),
	}
}

// Track will pick the least used tablet from "stats", increment its usage by 1
// and return it.
// "stats" must not be empty.
func (t *TabletTracker) Track(stats []discovery.TabletStats) *topodatapb.Tablet {
	if len(stats) == 0 {
		panic("stats must not be empty")
	}
	t.mu.Lock()
	defer t.mu.Unlock()

	// Try to find a tablet which is not in use yet.
	for _, stat := range stats {
		key := topoproto.TabletAliasString(stat.Tablet.Alias)
		if _, ok := t.usedTablets[key]; !ok {
			t.usedTablets[key] = 1
			return stat.Tablet
		}
	}

	// If we reached this point, "stats" is a subset of "usedTablets" i.e.
	// all tablets are already in use. Take the least used one.
	for _, aliasString := range t.tabletsByUsage() {
		for _, stat := range stats {
			key := topoproto.TabletAliasString(stat.Tablet.Alias)
			if key == aliasString {
				t.usedTablets[key]++
				return stat.Tablet
			}
		}
	}
	panic("BUG: we did not add any tablet")
}

// Untrack decrements the usage of "alias" by 1.
func (t *TabletTracker) Untrack(alias *topodatapb.TabletAlias) {
	t.mu.Lock()
	defer t.mu.Unlock()

	key := topoproto.TabletAliasString(alias)
	count, ok := t.usedTablets[key]
	if !ok {
		panic(fmt.Sprintf("tablet: %v was never tracked", key))
	}
	count--
	if count == 0 {
		delete(t.usedTablets, key)
	} else {
		t.usedTablets[key] = count
	}
}

// TabletsInUse returns a string of all tablet aliases currently in use.
// The tablets are separated by a space.
func (t *TabletTracker) TabletsInUse() string {
	t.mu.Lock()
	defer t.mu.Unlock()

	var aliases []string
	for alias := range t.usedTablets {
		aliases = append(aliases, alias)
	}
	sort.Strings(aliases)
	return strings.Join(aliases, " ")
}

func (t *TabletTracker) tabletsByUsage() []string {
	sorted := sortMapByValue(t.usedTablets)
	var tablets []string
	for _, pair := range sorted {
		tablets = append(tablets, pair.Key)
	}
	return tablets
}

// Sort by value was originally written by Andrew Gerrand:
// Source: https://groups.google.com/d/msg/golang-nuts/FT7cjmcL7gw/Gj4_aEsE_IsJ

// Pair represents a tablet (Key) and its usage (Value).
type Pair struct {
	Key   string
	Value int
}

// PairList is a slice of Pairs that implements sort.Interface to sort by Value.
type PairList []Pair

func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

// A function to turn a map into a PairList, then sort and return it.
func sortMapByValue(m map[string]int) PairList {
	p := make(PairList, len(m))
	i := 0
	for k, v := range m {
		p[i] = Pair{k, v}
		i++
	}
	sort.Sort(p)
	return p
}
