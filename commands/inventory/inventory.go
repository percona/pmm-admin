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

// Package inventory provides inventory commands.
package inventory

import (
	"gopkg.in/alecthomas/kingpin.v2"
)

// Hide inventory commands from help by default.
const hide = true

// register commands
var (
	inventoryC       = kingpin.Command("inventory", "Inventory commands").Hide(hide)
	inventoryListC   = inventoryC.Command("list", "List inventory commands").Hide(hide)
	inventoryAddC    = inventoryC.Command("add", "Add to inventory commands").Hide(hide)
	inventoryRemoveC = inventoryC.Command("remove", "Remove from inventory commands").Hide(hide)
	// Add command gobal flags
	AddAddressFlag     = inventoryAddC.Flag("host", "Hostname").String()
	AddPortFlag        = inventoryAddC.Flag("port", "Port number").Int64()
	AddServiceNameFlag = inventoryAddC.Flag("service-name", "Service name").String()
)

type getter interface {
	GetAddress() string
	GetServiceName() string
	GetPort() int64
}

// Types implementing the getter interface:
// - addServiceMySQLCommand
// - addServiceProxySQLCommand
// - addServiceMongoDBCommand
// - ddServicePostgreSQLCommand
func processGlobalAddFlags(cmd getter) (string, string, int64) {
	serviceName := cmd.GetServiceName()
	if *AddServiceNameFlag != "" {
		serviceName = *AddServiceNameFlag
	}

	address := cmd.GetAddress()
	if *AddAddressFlag != "" {
		address = *AddAddressFlag
	}

	port := cmd.GetPort()
	if *AddPortFlag != 0 {
		port = *AddPortFlag
	}

	return serviceName, address, port
}
