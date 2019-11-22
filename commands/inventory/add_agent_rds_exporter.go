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

var addAgentRDSExporterResultT = commands.ParseTemplate(`
RDS Exporter added.
Agent ID              : {{ .Agent.AgentID }}
PMM-Agent ID          : {{ .Agent.PMMAgentID }}
Service ID            : {{ .Agent.ServiceID }}
Listen port           : {{ .Agent.ListenPort }}

Status                : {{ .Agent.Status }}
Disabled              : {{ .Agent.Disabled }}
Custom labels         : {{ .Agent.CustomLabels }}
`)

type addAgentRDSExporterResult struct {
	Agent *agents.AddRDSExporterOKBodyRDSExporter `json:"proxysql_exporter"`
}

func (res *addAgentRDSExporterResult) Result() {}

func (res *addAgentRDSExporterResult) String() string {
	return commands.RenderTemplate(addAgentRDSExporterResultT, res)
}

type addAgentRDSExporterCommand struct {
	PMMAgentID          string
	ServiceID           string
	CustomLabels        string
	AWSAccessKey        string
	AWSSecretKey        string
	SkipConnectionCheck bool
}

func (cmd *addAgentRDSExporterCommand) Run() (commands.Result, error) {
	customLabels, err := commands.ParseCustomLabels(cmd.CustomLabels)
	if err != nil {
		return nil, err
	}

	params := &agents.AddRDSExporterParams{
		Body: agents.AddRDSExporterBody{
			PMMAgentID:          cmd.PMMAgentID,
			ServiceID:           cmd.ServiceID,
			CustomLabels:        customLabels,
			SkipConnectionCheck: cmd.SkipConnectionCheck,
			AWSAccessKey:        cmd.AWSAccessKey,
			AWSSecretKey:        cmd.AWSSecretKey,
		},
		Context: commands.Ctx,
	}

	resp, err := client.Default.Agents.AddRDSExporter(params)
	if err != nil {
		return nil, err
	}
	return &addAgentRDSExporterResult{
		Agent: resp.Payload.RDSExporter,
	}, nil
}

// register command
var (
	AddAgentRDSExporter  = new(addAgentRDSExporterCommand)
	AddAgentRDSExporterC = addAgentC.Command("rds-exporter", "Add rds_exporter to inventory").Hide(hide)
)

func init() {
	AddAgentRDSExporterC.Arg("pmm-agent-id", "The pmm-agent identifier which runs this instance").StringVar(&AddAgentRDSExporter.PMMAgentID)
	AddAgentRDSExporterC.Arg("service-id", "Service identifier").StringVar(&AddAgentRDSExporter.ServiceID)
	AddAgentRDSExporterC.Flag("custom-labels", "Custom user-assigned labels").StringVar(&AddAgentRDSExporter.CustomLabels)
	AddAgentRDSExporterC.Flag("skip-connection-check", "Skip connection check").BoolVar(&AddAgentRDSExporter.SkipConnectionCheck)
	AddAgentRDSExporterC.Flag("aws-access-key", "AWS Access Key ID").StringVar(&AddAgentRDSExporter.AWSAccessKey)
	AddAgentRDSExporterC.Flag("aws-secret-key", "AWS Secret Access Key").StringVar(&AddAgentRDSExporter.AWSSecretKey)
}
