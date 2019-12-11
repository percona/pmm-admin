// pmm-admin
// Copyright 2019 Percona LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package management

import (
	"gopkg.in/alecthomas/kingpin.v2"
)

// register command
var (
	AddC        = kingpin.Command("add", "Add Service to monitoring")
	Host        = AddC.Flag("host", "Hostname").String()
	Port        = AddC.Flag("port", "Port number").Int()
	ServiceName = AddC.Flag("service-name", "Service name").String()
)
