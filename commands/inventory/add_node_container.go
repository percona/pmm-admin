// pmm-admin
// Copyright (C) 2018 Percona LLC
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.
package inventory

import (
	"github.com/percona/pmm/api/inventory/json/client"
	"github.com/percona/pmm/api/inventory/json/client/nodes"

	"github.com/percona/pmm-admin/commands"
)

var addContainerNodeResultT = commands.ParseTemplate(`
Container Node added.
Node ID  : {{ .Node.NodeID }}
Node name: {{ .Node.NodeName }}

MachineID          : {{ .Node.MachineID }}
DockerContainerID  : {{ .Node.DockerContainerID }}
DockerContainerName: {{ .Node.DockerContainerName }}
`)

type addContainerNodeResult struct {
	Node *nodes.AddContainerNodeOKBodyContainer `json:"container"`
}

func (res *addContainerNodeResult) Result() {}

func (res *addContainerNodeResult) String() string {
	return commands.RenderTemplate(addContainerNodeResultT, res)
}

type addNodeContainerCommand struct {
	NodeType            string
	NodeName            string
	MachineID           string
	DockerContainerID   string
	DockerContainerName string
	Address             string
	CustomLabels        map[string]string
}

func (cmd *addNodeContainerCommand) Run() (commands.Result, error) {
	params := &nodes.AddContainerNodeParams{
		Body: nodes.AddContainerNodeBody{
			NodeName:            cmd.NodeName,
			MachineID:           cmd.MachineID,
			DockerContainerID:   cmd.DockerContainerID,
			DockerContainerName: cmd.DockerContainerName,
		},
		Context: commands.Ctx,
	}

	resp, err := client.Default.Nodes.AddContainerNode(params)
	if err != nil {
		return nil, err
	}
	return &addContainerNodeResult{
		Node: resp.Payload.Container,
	}, nil
}

// register command
var (
	AddNodeContainer  = new(addNodeContainerCommand)
	AddNodeContainerC = InventoryAddC.Command("container", "Add container node to inventory.")
)

func init() {
	AddNodeContainerC.Arg("name", "Node name").StringVar(&AddNodeContainer.NodeName)

	AddNodeContainerC.Flag("machine-id", "Linux machine-id.").StringVar(&AddNodeContainer.MachineID)
	AddNodeContainerC.Flag("container-id", "Docker container identifier. If specified, must be a unique Docker container identifier.").StringVar(&AddNodeContainer.DockerContainerID)
	AddNodeContainerC.Flag("container-name", "Container name.").StringVar(&AddNodeContainer.DockerContainerName)
	AddNodeContainerC.Flag("address", "Address.").StringVar(&AddNodeContainer.Address)
}
