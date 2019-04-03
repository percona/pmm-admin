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

package commands

import (
	"github.com/percona/pmm/api/managementpb/json/client"
	"github.com/percona/pmm/api/managementpb/json/client/node"
)

type RegisterResult struct {
}

func (res *RegisterResult) result() {}

func (res *RegisterResult) String() string {
	return ""
}

type Register struct {
	NodeType      string
	NodeName      string
	MachineID     string
	Distro        string
	ContainerID   string
	ContainerName string
}

func (cmd *Register) Run() (Result, error) {
	params := &node.RegisterParams{
		Body: node.RegisterBody{
			NodeName:      cmd.NodeName,
			NodeType:      &cmd.NodeType,
			MachineID:     cmd.MachineID,
			Distro:        cmd.Distro,
			ContainerID:   cmd.ContainerID,
			ContainerName: cmd.ContainerName,
		},
		Context: Ctx,
	}
	resp, err := client.Default.Node.Register(params)
	if err != nil {
		return nil, err
	}
	_ = resp
	return nil, nil
}

// check interfaces
var (
	_ Result  = (*RegisterResult)(nil)
	_ Command = (*Register)(nil)
)
