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

// Package services exposes all the services for the vtgateclienttest binary.
package services

import "github.com/mushiyu/vitess/go/vt/vtgate/vtgateservice"

// CreateServices creates the implementation chain of all the test cases
func CreateServices() vtgateservice.VTGateService {
	var s vtgateservice.VTGateService
	s = newTerminalClient()
	s = newSuccessClient(s)
	s = newErrorClient(s)
	s = newCallerIDClient(s)
	s = newEchoClient(s)
	return s
}
