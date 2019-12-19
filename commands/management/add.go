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
	"net"
	"strconv"

	"gopkg.in/alecthomas/kingpin.v2"
)

// register command
var (
	AddC = kingpin.Command("add", "Add Service to monitoring")

	// Add command global flags
	addAddressFlag     = AddC.Flag("host", "Hostname").String()
	addPortFlag        = AddC.Flag("port", "Port number").Int64()
	addServiceNameFlag = AddC.Flag("service-name", "Service name").String()
)

type getter interface {
	GetAddressPort() string
	GetServiceName() string
}

// Types implementing the getter interface:
// - addMongoDBCommand
// - addMySQLCommand
// - addPostgreSQLCommand
// - addProxySQLCommand
// Returns, service name, address, port, error
func processGlobalAddFlags(cmd getter) (string, string, int, error) {
	serviceName := cmd.GetServiceName()
	if *addServiceNameFlag != "" {
		serviceName = *addServiceNameFlag
	}

	addressPort := cmd.GetAddressPort()

	host, portS, err := net.SplitHostPort(addressPort)
	if err != nil {
		return "", "", 0, err
	}

	port, err := strconv.Atoi(portS)
	if err != nil {
		return "", "", 0, err
	}

	if *addAddressFlag != "" {
		host = *addAddressFlag
	}

	if *addPortFlag != 0 {
		port = int(*addPortFlag)
	}

	return serviceName, host, port, nil
}
