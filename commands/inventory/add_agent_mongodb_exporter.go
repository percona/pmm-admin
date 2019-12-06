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

var addAgentMongodbExporterResultT = commands.ParseTemplate(`
MongoDB Exporter added.
Agent ID              : {{ .Agent.AgentID }}
PMM-Agent ID          : {{ .Agent.PMMAgentID }}
Service ID            : {{ .Agent.ServiceID }}
Username              : {{ .Agent.Username }}
Listen port           : {{ .Agent.ListenPort }}
TLS enabled           : {{ .Agent.TLS }}
Skip TLS verification : {{ .Agent.TLSSkipVerify }}

Status                : {{ .Agent.Status }}
Disabled              : {{ .Agent.Disabled }}
Custom labels         : {{ .Agent.CustomLabels }}
`)

type addAgentMongodbExporterResult struct {
	Agent *agents.AddMongoDBExporterOKBodyMongodbExporter `json:"mongodb_exporter"`
}

func (res *addAgentMongodbExporterResult) Result() {}

func (res *addAgentMongodbExporterResult) String() string {
	return commands.RenderTemplate(addAgentMongodbExporterResultT, res)
}

type addAgentMongodbExporterCommand struct {
	PMMAgentID          string
	ServiceID           string
	Username            string
	Password            string
	CustomLabels        string
	SkipConnectionCheck bool
	TLS                 bool
	TLSSkipVerify       bool
}

func (cmd *addAgentMongodbExporterCommand) Run() (commands.Result, error) {
	customLabels, err := commands.ParseCustomLabels(cmd.CustomLabels)
	if err != nil {
		return nil, err
	}
	params := &agents.AddMongoDBExporterParams{
		Body: agents.AddMongoDBExporterBody{
			PMMAgentID:          cmd.PMMAgentID,
			ServiceID:           cmd.ServiceID,
			Username:            cmd.Username,
			Password:            cmd.Password,
			CustomLabels:        customLabels,
			SkipConnectionCheck: cmd.SkipConnectionCheck,
			TLS:                 cmd.TLS,
			TLSSkipVerify:       cmd.TLSSkipVerify,
		},
		Context: commands.Ctx,
	}

	resp, err := client.Default.Agents.AddMongoDBExporter(params)
	if err != nil {
		return nil, err
	}
	return &addAgentMongodbExporterResult{
		Agent: resp.Payload.MongodbExporter,
	}, nil
}

// register command
var (
	AddAgentMongodbExporter  = new(addAgentMongodbExporterCommand)
	AddAgentMongodbExporterC = addAgentC.Command("mongodb-exporter", "Add mongodb_exporter to inventory").Hide(hide)
)

func init() {
	AddAgentMongodbExporterC.Arg("pmm-agent-id", "The pmm-agent identifier which runs this instance").Required().StringVar(&AddAgentMongodbExporter.PMMAgentID)
	AddAgentMongodbExporterC.Arg("service-id", "Service identifier").Required().StringVar(&AddAgentMongodbExporter.ServiceID)
	AddAgentMongodbExporterC.Arg("username", "MongoDB username for scraping metrics").StringVar(&AddAgentMongodbExporter.Username)
	AddAgentMongodbExporterC.Flag("password", "MongoDB password for scraping metrics").StringVar(&AddAgentMongodbExporter.Password)
	AddAgentMongodbExporterC.Flag("custom-labels", "Custom user-assigned labels").StringVar(&AddAgentMongodbExporter.CustomLabels)
	AddAgentMongodbExporterC.Flag("skip-connection-check", "Skip connection check").BoolVar(&AddAgentMongodbExporter.SkipConnectionCheck)
	AddAgentMongodbExporterC.Flag("tls", "Use TLS to connect to the database").BoolVar(&AddAgentMongodbExporter.TLS)
	AddAgentMongodbExporterC.Flag("tls-skip-verify", "Skip TLS certificates validation").BoolVar(&AddAgentMongodbExporter.TLSSkipVerify)
}
