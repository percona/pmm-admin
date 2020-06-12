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

package commands

import (
	"github.com/percona/pmm/api/inventorypb/json/client"
	"github.com/percona/pmm/api/inventorypb/json/client/services"

	"github.com/percona/pmm-admin/agentlocal"
)

func (cmd *annotationCommand) serviceNames() ([]string, error) {
	var servicesNameList []string
	if cmd.Service {
		status, err := agentlocal.GetStatus(agentlocal.DoNotRequestNetworkInfo)
		if err != nil {
			return nil, err
		}

		params := &services.ListServicesParams{
			Body: services.ListServicesBody{
				NodeID: status.NodeID,
			},
			Context: Ctx,
		}

		result, err := client.Default.Services.ListServices(params)
		if err != nil {
			return nil, err
		}

		for _, s := range result.Payload.Mysql {
			servicesNameList = append(servicesNameList, s.ServiceName)
		}
		for _, s := range result.Payload.Mongodb {
			servicesNameList = append(servicesNameList, s.ServiceName)
		}
		for _, s := range result.Payload.Postgresql {
			servicesNameList = append(servicesNameList, s.ServiceName)
		}
		for _, s := range result.Payload.Proxysql {
			servicesNameList = append(servicesNameList, s.ServiceName)
		}
		for _, s := range result.Payload.External {
			servicesNameList = append(servicesNameList, s.ServiceName)
		}
	}

	if cmd.ServiceName != "" {
		servicesNameList = []string{cmd.ServiceName}
	}

	return servicesNameList, nil
}
