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

package vtexplain

import (
	"encoding/json"
	"testing"

	topodatapb "github.com/mushiyu/vitess/go/vt/proto/topodata"
)

func TestParseSchema(t *testing.T) {
	testSchema := `
create table t1 (
	id bigint(20) unsigned not null default 123,
	val varchar default "default",
	primary key (id)
);

create table t2 (
	val text default "default2"
);

create table t3 (
    b bit(1) default B'0'
);

create table t4 like t3;

create table t5 (like t2);
`

	ddls, err := parseSchema(testSchema, &Options{StrictDDL: true})
	if err != nil {
		t.Fatalf("parseSchema: %v", err)
	}
	initTabletEnvironment(ddls, defaultTestOpts())

	tablet := newTablet(defaultTestOpts(), &topodatapb.Tablet{
		Keyspace: "test_keyspace",
		Shard:    "-80",
	})
	se := tablet.tsv.SchemaEngine()
	tables := se.GetSchema()

	t1 := tables["t1"]
	if t1 == nil {
		t.Fatalf("table t1 wasn't parsed properly")
	}

	wantCols := `[{"Name":"id","Type":778,"IsAuto":false,"Default":123},{"Name":"val","Type":6165,"IsAuto":false,"Default":"'default'"}]`
	got, _ := json.Marshal(t1.Columns)
	if wantCols != string(got) {
		t.Errorf("expected %s got %s", wantCols, string(got))
	}

	if !t1.HasPrimary() || len(t1.PKColumns) != 1 || t1.PKColumns[0] != 0 {
		t.Errorf("expected HasPrimary && t1.PKColumns == [0] got %v", t1.PKColumns)
	}
	pkCol := t1.GetPKColumn(0)
	if pkCol == nil || pkCol.String() != "{Name: 'id', Type: UINT64}" {
		t.Errorf("expected pkCol[0] == id, got %v", pkCol)
	}

	t2 := tables["t2"]
	if t2 == nil {
		t.Fatalf("table t2 wasn't parsed properly")
	}

	wantCols = `[{"Name":"val","Type":6163,"IsAuto":false,"Default":"'default2'"}]`
	got, _ = json.Marshal(t2.Columns)
	if wantCols != string(got) {
		t.Errorf("expected %s got %s", wantCols, string(got))
	}

	if t2.HasPrimary() || len(t2.PKColumns) != 0 {
		t.Errorf("expected !HasPrimary && t2.PKColumns == [] got %v", t2.PKColumns)
	}

	t5 := tables["t5"]
	if t5 == nil {
		t.Fatalf("table t5 wasn't parsed properly")
	}
	got, _ = json.Marshal(t5.Columns)
	if wantCols != string(got) {
		t.Errorf("expected %s got %s", wantCols, string(got))
	}

	if t5.HasPrimary() || len(t5.PKColumns) != 0 {
		t.Errorf("expected !HasPrimary && t5.PKColumns == [] got %v", t5.PKColumns)
	}
}

func TestErrParseSchema(t *testing.T) {
	testSchema := `
create table t1 like t2;
`
	expected := "check your schema, table[t2] doesnt exist"
	ddl, err := parseSchema(testSchema, &Options{StrictDDL: true})
	if err != nil {
		t.Fatalf("parseSchema: %v", err)
	}
	err = initTabletEnvironment(ddl, defaultTestOpts())
	if err.Error() != expected {
		t.Errorf("want: %s, got %s", expected, err.Error())
	}
}
