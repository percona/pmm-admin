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

package inventory

import (
	"github.com/percona/pmm/api/inventorypb/json/client"
	"github.com/percona/pmm/api/inventorypb/json/client/nodes"

	"github.com/percona/pmm-admin/commands"
)

var addRemoteRDSNodeResultT = commands.ParseTemplate(`
RDS Node added.
Node ID  : {{ .Node.NodeID }}
Node name: {{ .Node.NodeName }}

Address       : {{ .Node.Address }}
Custom labels : {{ .Node.CustomLabels }}

Region    : {{ .Node.Region }}
Az        : {{ .Node.Az }}
`)

type addRemoteRDSNodeResult struct {
	Node *nodes.AddRemoteRDSNodeOKBodyRemote `json:"remote"`
}

func (res *addRemoteRDSNodeResult) Result() {}

func (res *addRemoteRDSNodeResult) String() string {
	return commands.RenderTemplate(addNodeRemoteResultT, res)
}

type addRemoteRDSNodeCommand struct {
	NodeName     string
	Address      string
	CustomLabels string
	Region       string
	Az           string
}

func (cmd *addRemoteRDSNodeCommand) Run() (commands.Result, error) {
	customLabels, err := commands.ParseCustomLabels(cmd.CustomLabels)
	if err != nil {
		return nil, err
	}
	params := &nodes.AddRemoteRDSNodeParams{
		Body: nodes.AddRemoteRDSNodeBody{
			NodeName:     cmd.NodeName,
			Address:      cmd.Address,
			CustomLabels: customLabels,

			Region: cmd.Region,
			Az:     cmd.Az,
		},
		Context: commands.Ctx,
	}

	resp, err := client.Default.Nodes.AddRemoteRDSNode(params)
	if err != nil {
		return nil, err
	}
	return &addRemoteRDSNodeResult{
		Node: resp.Payload.Remote,
	}, nil
}

// register command
var (
	AddRemoteRDSNode  = new(addRemoteRDSNodeCommand)
	AddRemoteRDSNodeC = addNodeC.Command("rds-node", "Add RDS node to inventory")
)

func init() {
	AddRemoteRDSNodeC.Arg("name", "Node name").StringVar(&AddRemoteRDSNode.NodeName)

	AddRemoteRDSNodeC.Flag("address", "Address").StringVar(&AddRemoteRDSNode.Address)
	AddRemoteRDSNodeC.Flag("custom-labels", "Custom user-assigned labels").StringVar(&AddRemoteRDSNode.CustomLabels)

	AddRemoteRDSNodeC.Flag("region", "Node region").StringVar(&AddRemoteRDSNode.Region)
	AddRemoteRDSNodeC.Flag("az", "Node availability zone").StringVar(&AddRemoteRDSNode.Az)
}
