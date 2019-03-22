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
	"github.com/percona/pmm/api/inventory/json/client/agents"
	"github.com/percona/pmm/api/inventory/json/client/nodes"
	"github.com/percona/pmm/api/inventory/json/client/services"
)

type ListResult struct {
	Nodes    []nodes.ListNodesOKBody
	Services []services.ListServicesOKBody
	Agents   []agents.ListAgentsOKBody
}

func (r *ListResult) result() {}

func (r *ListResult) String() string {
	return ""
}

type ListCmd struct {
	CommonParams
}

func (cmd *ListCmd) Run() Result {
	return &ListResult{}
}

// check interfaces
var (
	_ Result = (*ListResult)(nil)
	_ Cmd    = (*ListCmd)(nil)
)
