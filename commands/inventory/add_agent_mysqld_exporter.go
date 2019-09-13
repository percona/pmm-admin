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

var addAgentMysqldExporterResultT = commands.ParseTemplate(`
Mysqld Exporter added.
Agent ID              : {{ .Agent.AgentID }}
PMM-Agent ID          : {{ .Agent.PMMAgentID }}
Service ID            : {{ .Agent.ServiceID }}
Username              : {{ .Agent.Username }}
Password              : {{ .Agent.Password }}
Listen port           : {{ .Agent.ListenPort }}
TLS enabled           : {{ .Agent.TLS }}
Skip TLS verification : {{ .Agent.TLSSkipVerify }}

Status                : {{ .Agent.Status }}
Disabled              : {{ .Agent.Disabled }}
Custom labels         : {{ .Agent.CustomLabels }}
`)

type addAgentMysqldExporterResult struct {
	Agent *agents.AddMySQLdExporterOKBodyMysqldExporter `json:"mysqld_exporter"`
}

func (res *addAgentMysqldExporterResult) Result() {}

func (res *addAgentMysqldExporterResult) String() string {
	return commands.RenderTemplate(addAgentMysqldExporterResultT, res)
}

type addAgentMysqldExporterCommand struct {
	PMMAgentID          string
	ServiceID           string
	Username            string
	Password            string
	CustomLabels        string
	SkipConnectionCheck bool
	TLS                 bool
	TLSSkipVerify       bool
}

func (cmd *addAgentMysqldExporterCommand) Run() (commands.Result, error) {
	customLabels, err := commands.ParseCustomLabels(cmd.CustomLabels)
	if err != nil {
		return nil, err
	}
	params := &agents.AddMySQLdExporterParams{
		Body: agents.AddMySQLdExporterBody{
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

	resp, err := client.Default.Agents.AddMySQLdExporter(params)
	if err != nil {
		return nil, err
	}
	return &addAgentMysqldExporterResult{
		Agent: resp.Payload.MysqldExporter,
	}, nil
}

// register command
var (
	AddAgentMysqldExporter  = new(addAgentMysqldExporterCommand)
	AddAgentMysqldExporterC = addAgentC.Command("mysqld-exporter", "Add mysqld_exporter to inventory").Hide(hide)
)

func init() {
	AddAgentMysqldExporterC.Arg("pmm-agent-id", "The pmm-agent identifier which runs this instance").StringVar(&AddAgentMysqldExporter.PMMAgentID)
	AddAgentMysqldExporterC.Arg("service-id", "Service identifier").StringVar(&AddAgentMysqldExporter.ServiceID)
	AddAgentMysqldExporterC.Arg("username", "MySQL username for scraping metrics").Default("root").StringVar(&AddAgentMysqldExporter.Username)
	AddAgentMysqldExporterC.Flag("password", "MySQL password for scraping metrics").StringVar(&AddAgentMysqldExporter.Password)
	AddAgentMysqldExporterC.Flag("custom-labels", "Custom user-assigned labels").StringVar(&AddAgentMysqldExporter.CustomLabels)
	AddAgentMysqldExporterC.Flag("skip-connection-check", "Skip connection check").BoolVar(&AddAgentMysqldExporter.SkipConnectionCheck)
	AddAgentMysqldExporterC.Flag("tls", "Use TLS to connect to the database").BoolVar(&AddAgentMysqldExporter.TLS)
	AddAgentMysqldExporterC.Flag("tls-skip-verify", "Skip TLS certificates validation").BoolVar(&AddAgentMysqldExporter.TLSSkipVerify)
}
