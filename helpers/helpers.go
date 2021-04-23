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

// Package helpers provides helpers for whole pmm-admin.
package helpers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/percona/pmm-admin/agentlocal"
)

// GetServerVersion return version of PMM Server.
func GetServerVersion() (float64, error) {
	status, err := agentlocal.GetStatus(agentlocal.DoNotRequestNetworkInfo)
	if err != nil {
		return 0, err
	}

	split := strings.Split(status.ServerVersion, "-")
	split = strings.Split(split[0], ".")
	if len(split) < 3 {
		return 0, fmt.Errorf("failed to parse server version %s", status.ServerVersion)
	}

	f, err := strconv.ParseFloat(fmt.Sprintf("%s.%s%s", split[0], split[1], split[2]), 64)
	if err != nil {
		return 0, err
	}

	return f, nil
}
