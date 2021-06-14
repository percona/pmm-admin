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

package management

import (
	"github.com/percona/pmm/api/inventorypb/json/client"
	"github.com/percona/pmm/api/inventorypb/json/client/nodes"
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/percona/pmm-admin/agentlocal"
	"github.com/percona/pmm-admin/commands"
)

type unregisterCommand struct {
	Force bool
}
type unregisterResult struct{}

var unregisterNodeResultT = commands.ParseTemplate(`
Local node unregistered.
`)

func (res *unregisterResult) Result() {}

func (res *unregisterResult) String() string {
	return commands.RenderTemplate(unregisterNodeResultT, res)
}

func (cmd *unregisterCommand) Run() (commands.Result, error) {
	status, err := agentlocal.GetStatus(agentlocal.DoNotRequestNetworkInfo)
	if err != nil {
		return nil, err
	}

	params := &nodes.RemoveNodeParams{
		Body: nodes.RemoveNodeBody{
			NodeID: status.NodeID,
			Force:  cmd.Force,
		},
		Context: commands.Ctx,
	}

	_, err = client.Default.Nodes.RemoveNode(params)
	if err != nil {
		return nil, err
	}

	return new(unregisterResult), nil
}

// unregister command
var (
	Unregister  = new(unregisterCommand)
	UnregisterC = kingpin.Command("unregister", "Unregister current Node from PMM Server")
)

func init() {
	UnregisterC.Flag("force", "Remove this node with all dependencies").BoolVar(&Unregister.Force)
}
