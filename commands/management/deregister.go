// pmm-admin
// Copyright 2021 Percona LLC
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
	"github.com/percona/pmm-admin/agentlocal"
	"github.com/percona/pmm-admin/commands"
	"github.com/percona/pmm/api/inventorypb/json/client"
	"github.com/percona/pmm/api/inventorypb/json/client/nodes"
	"gopkg.in/alecthomas/kingpin.v2"
)

type deregisterCommand struct {
	Force bool
}
type deregisterResult struct {}

var deregisterNodeResultT = commands.ParseTemplate(`
Local node deregistered.
`)

func (res *deregisterResult) Result() {}

func (res *deregisterResult) String() string {
	return commands.RenderTemplate(deregisterNodeResultT, res)
}

func (cmd *deregisterCommand) Run() (commands.Result, error) {
	agentlocalstatus, query_err := agentlocal.GetStatus(agentlocal.RequestNetworkInfo)
	if query_err != nil {
		return nil, query_err
	}
	node_id := agentlocalstatus.NodeID

	params := &nodes.RemoveNodeParams{
		Body: nodes.RemoveNodeBody{
			NodeID: node_id,
			Force: cmd.Force,
		},
		Context: commands.Ctx,
	}

	_, client_err := client.Default.Nodes.RemoveNode(params)
	if client_err != nil {
		return nil, client_err
	}

	return new(deregisterResult), nil
}

// deregister command
var (
	Deregister = new(deregisterCommand)
	DeregisterC = kingpin.Command("deregister", "Register current Node at PMM Server")
)

func init() {
	DeregisterC.Flag("force", "Remove thisnode with all dependencies").BoolVar(&Deregister.Force)
}
