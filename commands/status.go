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

package commands

import (
	"fmt"
	"os"
	"strings"
	"time"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/percona/pmm-admin/agentlocal"
)

var summaryResultT = ParseTemplate(`
Agent ID: {{ .PMMAgentStatus.AgentID }}
Node ID : {{ .PMMAgentStatus.NodeID }}

PMM Server:
	URL    : {{ .PMMAgentStatus.ServerURL }}
	Version: {{ .PMMAgentStatus.ServerVersion }}

PMM-agent:
	Connected : {{ .PMMAgentStatus.Connected }}{{ if .PMMAgentStatus.Connected }}
	Time drift: {{ .PMMAgentStatus.ServerClockDrift }}
	Latency   : {{ .PMMAgentStatus.ServerLatency }}
{{ end }}
Agents:
{{ range .PMMAgentStatus.Agents }}	{{ .AgentID }} {{ .AgentType }} {{ .Status }}
{{ end }}
`)

type summaryResult struct {
	PMMAgentStatus *agentlocal.Status `json:"pmm_agent_status"`
}

func (res *summaryResult) Result() {}

func (res *summaryResult) String() string {
	return RenderTemplate(summaryResultT, res)
}

type statusCommand struct {
	Archive         bool
	ArchiveFilename string
}

func (cmd *statusCommand) Run() (Result, error) {
	status, err := agentlocal.GetStatus(agentlocal.RequestNetworkInfo)
	if err != nil {
		return nil, err
	}

	res := &summaryResult{
		PMMAgentStatus: status,
	}
	if !cmd.Archive {
		return res, nil
	}

	return res, nil
}

// register command
var (
	Summary  = new(statusCommand)
	SummaryC = kingpin.Command("summary", "Show summary status information")
	StatusC  = kingpin.Command("status", "").Hidden() // TODO remove it https://jira.percona.com/browse/PMM-4704
)

func init() {
	SummaryC.Flag("archive", "Generate summary archive file").BoolVar(&Summary.Archive)

	hostname, _ := os.Hostname()
	archiveFilename := fmt.Sprintf("summary_%s_%s.zip",
		strings.Replace(hostname, ".", "_", -1), time.Now().Format("2006_01_02_15_04_05"))
	SummaryC.Flag("archive-file", "Summary archive filename").Default(archiveFilename).StringVar(&Summary.ArchiveFilename)
}
