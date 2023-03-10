/*
Copyright 2018 The Vitess Authors

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

// Package vtaclcheck analyzes a set of sql statements and returns the
// corresponding vtgate and vttablet query plans that will be executed
// on the given statements
package vtaclcheck

import (
	"fmt"

	"github.com/mushiyu/vitess/go/vt/tableacl"
	"github.com/mushiyu/vitess/go/vt/tableacl/simpleacl"
)

// Options to control the explain process
type Options struct {
	// AclFile is the file with the JSON acl configuration
	ACLFile string
}

var options *Options

// Init sets up the fake execution environment
func Init(opts *Options) error {
	// verify opts is defined
	if opts == nil {
		return fmt.Errorf("vtaclcheck.Init: opts is NULL")
	}
	// Verify options
	if len(opts.ACLFile) == 0 {
		return fmt.Errorf("-acl_file <filename> provided but len(filename) is zero")
	}

	options = opts

	return nil
}

// Run the check on the given file
func Run() error {
	filename := options.ACLFile

	tableacl.Register("simpleacl", &simpleacl.Factory{})
	err := tableacl.Init(
		filename,
		func() {},
	)
	if err != nil {
		return fmt.Errorf("Fail to initialize Table ACL: %v", err)
	}

	fmt.Printf("JSON ACL file %s looks good\n", filename)

	return nil
}
