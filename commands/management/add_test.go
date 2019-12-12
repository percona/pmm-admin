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
	"fmt"
	"testing"

	"github.com/AlekSi/pointer"
	"github.com/stretchr/testify/assert"
)

func TestGlobalFlags(t *testing.T) {
	wantHost := "localhost"
	wantPort := 27017
	cmd := &addMongoDBCommand{
		AddressPort: fmt.Sprintf("%s:%d", wantHost, wantPort),
		ServiceName: "mongodb-service",
	}

	serviceName, address, port, err := processGlobalAddFlags(cmd)

	assert.NoError(t, err)
	assert.Equal(t, serviceName, cmd.ServiceName)
	assert.Equal(t, address, wantHost)
	assert.Equal(t, port, wantPort)

	// Flags have precedence over positional arguments
	AddAddressFlag = pointer.ToString("remotehost")
	AddPortFlag = pointer.ToInt64(27018)
	AddServiceNameFlag = pointer.ToString("new-service")
	serviceName, address, port, err = processGlobalAddFlags(cmd)

	assert.NoError(t, err)
	assert.Equal(t, serviceName, *AddServiceNameFlag)
	assert.Equal(t, address, *AddAddressFlag)
	assert.Equal(t, int64(port), *AddPortFlag)
}
