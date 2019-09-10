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
	"github.com/percona/pmm/api/inventorypb/json/client/agents"

	"github.com/percona/pmm-admin/commands"
)

var addAgentQANMongoDBProfilerAgentResultT = commands.ParseTemplate(`
QAN MongoDB profiler agent added.
Agent ID     : {{ .Agent.AgentID }}
PMM-Agent ID : {{ .Agent.PMMAgentID }}
Service ID   : {{ .Agent.ServiceID }}
Username     : {{ .Agent.Username }}
Password     : {{ .Agent.Password }}

Status       : {{ .Agent.Status }}
Disabled     : {{ .Agent.Disabled }}
Custom labels: {{ .Agent.CustomLabels }}
`)

type addAgentQANMongoDBProfilerAgentResult struct {
	Agent *agents.AddQANMongoDBProfilerAgentOKBodyQANMongodbProfilerAgent `json:"qan_mongodb_profiler_agent"`
}

func (res *addAgentQANMongoDBProfilerAgentResult) Result() {}

func (res *addAgentQANMongoDBProfilerAgentResult) String() string {
	return commands.RenderTemplate(addAgentQANMongoDBProfilerAgentResultT, res)
}

type addAgentQANMongoDBProfilerAgentCommand struct {
	PMMAgentID          string
	ServiceID           string
	Username            string
	Password            string
	CustomLabels        string
	SkipConnectionCheck bool
}

func (cmd *addAgentQANMongoDBProfilerAgentCommand) Run() (commands.Result, error) {
	customLabels, err := commands.ParseCustomLabels(cmd.CustomLabels)
	if err != nil {
		return nil, err
	}
	params := &agents.AddQANMongoDBProfilerAgentParams{
		Body: agents.AddQANMongoDBProfilerAgentBody{
			PMMAgentID:          cmd.PMMAgentID,
			ServiceID:           cmd.ServiceID,
			Username:            cmd.Username,
			Password:            cmd.Password,
			CustomLabels:        customLabels,
			SkipConnectionCheck: cmd.SkipConnectionCheck,
		},
		Context: commands.Ctx,
	}

	resp, err := client.Default.Agents.AddQANMongoDBProfilerAgent(params)
	if err != nil {
		return nil, err
	}
	return &addAgentQANMongoDBProfilerAgentResult{
		Agent: resp.Payload.QANMongodbProfilerAgent,
	}, nil
}

// register command
var (
	AddAgentQANMongoDBProfilerAgent  = new(addAgentQANMongoDBProfilerAgentCommand)
	AddAgentQANMongoDBProfilerAgentC = addAgentC.Command("qan-mongodb-profiler-agent", "add QAN MongoDB profiler agent to inventory")
)

func init() {
	AddAgentQANMongoDBProfilerAgentC.Arg("pmm-agent-id", "The pmm-agent identifier which runs this instance").StringVar(&AddAgentQANMongoDBProfilerAgent.PMMAgentID)
	AddAgentQANMongoDBProfilerAgentC.Arg("service-id", "Service identifier").StringVar(&AddAgentQANMongoDBProfilerAgent.ServiceID)
	AddAgentQANMongoDBProfilerAgentC.Arg("username", "MongoDB username for scraping metrics").StringVar(&AddAgentQANMongoDBProfilerAgent.Username)
	AddAgentQANMongoDBProfilerAgentC.Flag("password", "MongoDB password for scraping metrics").StringVar(&AddAgentQANMongoDBProfilerAgent.Password)
	AddAgentQANMongoDBProfilerAgentC.Flag("custom-labels", "Custom user-assigned labels").StringVar(&AddAgentQANMongoDBProfilerAgent.CustomLabels)
	AddAgentQANMongoDBProfilerAgentC.Flag("skip-connection-check", "Skip connection check").BoolVar(&AddAgentQANMongoDBProfilerAgent.SkipConnectionCheck)
}
