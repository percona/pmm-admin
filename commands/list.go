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
	"github.com/percona/pmm/api/inventory/json/client"
	"github.com/percona/pmm/api/inventory/json/client/agents"
	"github.com/percona/pmm/api/inventory/json/client/nodes"
	"github.com/percona/pmm/api/inventory/json/client/services"
)

type ListResult struct {
	Nodes    *nodes.ListNodesOKBody
	Services *services.ListServicesOKBody
	Agents   *agents.ListAgentsOKBody
}

func (res *ListResult) result() {}

func (res *ListResult) String() string {
	return ""
}

type List struct {
}

func (cmd *List) Run() (Result, error) {
	nodes, err := client.Default.Nodes.ListNodes(nil)
	if err != nil {
		return nil, err
	}
	services, err := client.Default.Services.ListServices(nil)
	if err != nil {
		return nil, err
	}
	agents, err := client.Default.Agents.ListAgents(nil)
	if err != nil {
		return nil, err
	}
	return &ListResult{
		Nodes:    nodes.Payload,
		Services: services.Payload,
		Agents:   agents.Payload,
	}, nil
}

// check interfaces
var (
	_ Result  = (*ListResult)(nil)
	_ Command = (*List)(nil)
)
