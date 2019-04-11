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
	"github.com/percona/pmm/api/inventorypb/json/client"
	"github.com/percona/pmm/api/inventorypb/json/client/nodes"

	"github.com/percona/pmm-admin/commands"
)

var addNodeGenericResultT = commands.ParseTemplate(`
Generic Node added.
Node ID  : {{ .Node.NodeID }}
Node name: {{ .Node.NodeName }}

MachineID    : {{ .Node.MachineID }}
Distro       : {{ .Node.Distro }}
DistroVersion: {{ .Node.DistroVersion }}
Address      : {{ .Node.Address }}
CustomLabels : {{ .Node.CustomLabels }}
`)

type addNodeGenericResult struct {
	Node *nodes.AddGenericNodeOKBodyGeneric `json:"generic"`
}

func (res *addNodeGenericResult) Result() {}

func (res *addNodeGenericResult) String() string {
	return commands.RenderTemplate(addNodeGenericResultT, res)
}

type addNodeGenericCommand struct {
	NodeName      string
	MachineID     string
	Distro        string
	DistroVersion string
	Address       string
	CustomLabels  string
}

func (cmd *addNodeGenericCommand) Run() (commands.Result, error) {
	customLabels, err := parseCustomLabels(cmd.CustomLabels)
	if err != nil {
		return nil, err
	}
	params := &nodes.AddGenericNodeParams{
		Body: nodes.AddGenericNodeBody{
			NodeName:      cmd.NodeName,
			MachineID:     cmd.MachineID,
			Distro:        cmd.Distro,
			DistroVersion: cmd.DistroVersion,
			Address:       cmd.Address,
			CustomLabels:  customLabels,
		},
		Context: commands.Ctx,
	}

	resp, err := client.Default.Nodes.AddGenericNode(params)
	if err != nil {
		return nil, err
	}
	return &addNodeGenericResult{
		Node: resp.Payload.Generic,
	}, nil
}

// register command
var (
	AddNodeGeneric  = new(addNodeGenericCommand)
	AddNodeGenericC = addNodeC.Command("generic", "Add generic node to inventory.")
)

func init() {
	AddNodeGenericC.Arg("name", "Node name").StringVar(&AddNodeGeneric.NodeName)

	AddNodeGenericC.Flag("machine-id", "Linux machine-id.").StringVar(&AddNodeGeneric.MachineID)
	AddNodeGenericC.Flag("distro", "Linux distribution (if any).").StringVar(&AddNodeGeneric.Distro)
	AddNodeGenericC.Flag("distro-version", "Linux distribution version (if any).").StringVar(&AddNodeGeneric.DistroVersion)
	AddNodeGenericC.Flag("address", "Address.").StringVar(&AddNodeGeneric.Address)
	AddNodeGenericC.Flag("custom-labels", "Custom user-assigned labels.").StringVar(&AddNodeGeneric.CustomLabels)
}
