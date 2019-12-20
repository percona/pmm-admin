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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestManagementGlobalFlags(t *testing.T) {
	tests := []struct {
		testName       string
		addressPortArg string
		serviceNameArg string
		addressFlag    string
		portFlag       uint16
		serviceFlag    string
		//
		wantHost    string
		wantPort    uint16
		wantService string
	}{
		{
			testName:       "Only positional arguments",
			addressPortArg: "localhost:27017",
			serviceNameArg: "service-name",

			wantHost:    "localhost",
			wantPort:    27017,
			wantService: "service-name",
		},
		{
			testName:       "Override only host",
			addressPortArg: "localhost:27017",
			serviceNameArg: "service-name",

			addressFlag: "visitant-host",

			wantHost:    "visitant-host",
			wantPort:    27017,
			wantService: "service-name",
		},
		{
			testName:       "Override only port",
			addressPortArg: "localhost:27017",
			serviceNameArg: "service-name",

			portFlag: 27018,

			wantHost:    "localhost",
			wantPort:    27018,
			wantService: "service-name",
		},
		{
			testName:       "Override only service name",
			addressPortArg: "localhost:27017",
			serviceNameArg: "service-name",

			serviceFlag: "no-service",

			wantHost:    "localhost",
			wantPort:    27017,
			wantService: "no-service",
		},
		{
			testName:       "Override everything",
			addressPortArg: "localhost:27017",
			serviceNameArg: "service-name",

			addressFlag: "new-address",
			portFlag:    27019,
			serviceFlag: "out-of-service",

			wantHost:    "new-address",
			wantPort:    27019,
			wantService: "out-of-service",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.testName, func(t *testing.T) {
			cmd := &addMongoDBCommand{
				Address:     test.addressPortArg,
				ServiceName: test.serviceNameArg,
			}
			addHostFlag = &test.addressFlag
			addPortFlag = &test.portFlag
			addServiceNameFlag = &test.serviceFlag

			serviceName, address, port, err := processGlobalAddFlags(cmd)

			assert.NoError(t, err)
			assert.Equal(t, serviceName, test.wantService)
			assert.Equal(t, address, test.wantHost)
			assert.Equal(t, int(port), int(test.wantPort))
		})
	}
}
