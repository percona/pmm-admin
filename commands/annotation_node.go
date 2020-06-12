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
	"github.com/percona/pmm/api/inventorypb/json/client/nodes"

	"github.com/percona/pmm-admin/agentlocal"
)

func (cmd *annotationCommand) nodeName() (string, error) {
	nodeName := ""
	if cmd.Node {
		status, err := agentlocal.GetStatus(agentlocal.DoNotRequestNetworkInfo)
		if err != nil {
			return "", err
		}

		params := &nodes.GetNodeParams{
			Body: nodes.GetNodeBody{
				NodeID: status.NodeID,
			},
			Context: Ctx,
		}

		result, err := client.Default.Nodes.GetNode(params)
		if err != nil {
			return "", err
		}

		nodeName = result.Payload.Generic.NodeName
	}

	if cmd.NodeName != "" {
		nodeName = cmd.NodeName
	}

	return nodeName, nil
}
