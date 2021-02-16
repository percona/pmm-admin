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
	"bytes"
	"fmt"
	"text/tabwriter"

	"github.com/percona/pmm/api/inventorypb/json/client"
	"github.com/percona/pmm/api/inventorypb/json/client/tunnels"

	"github.com/percona/pmm-admin/commands"
)

var listTunnelsResultT = commands.ParseTemplate(`
Agents list.

{{ printf "%-27s" "Agent type" }} {{ printf "%-15s" "Status" }} {{ printf "%-47s" "Agent ID" }} {{ printf "%-47s" "PMM-Agent ID" }} {{ printf "%-47s" "Service ID" }}
{{ range .Agents }}
{{- printf "%-27s" .HumanReadableAgentType }} {{ printf "%-15s" .NiceAgentStatus }} {{ .AgentID }}  {{ .PMMAgentID }}  {{ .ServiceID }}
{{ end }}
`)

type listResultTunnel struct {
	TunnelID       string `json:"tunnel_id"`
	ListenAgentID  string `json:"listen_agent_id"`
	ListenPort     uint16 `json:"listen_port"`
	ConnectAgentID string `json:"connect_agent_id"`
	ConnectPort    uint16 `json:"connect_port"`
}

type listTunnelsResult struct {
	Tunnels []listResultTunnel `json:"tunnels"`
}

func (res *listTunnelsResult) Result() {}

func (res *listTunnelsResult) String() string {
	if len(res.Tunnels) == 0 {
		return "No tunnels."
	}

	var buf bytes.Buffer
	buf.WriteString("Tunnels list.\n")

	w := tabwriter.NewWriter(&buf, 0, 0, 1, ' ', 0)
	fmt.Fprintf(w, "ID\tInfo\n")
	for _, t := range res.Tunnels {
		fmt.Fprintf(w, "%s\tlala\n", t.TunnelID)
	}
	w.Flush()
	return buf.String()
}

type listTunnelsCommand struct {
	filters tunnels.ListTunnelsBody
}

func (cmd *listTunnelsCommand) Run() (commands.Result, error) {
	params := &tunnels.ListTunnelsParams{
		Body:    cmd.filters,
		Context: commands.Ctx,
	}
	tunnelsRes, err := client.Default.Tunnels.ListTunnels(params)
	if err != nil {
		return nil, err
	}

	tunnelsList := make([]listResultTunnel, len(tunnelsRes.Payload.Tunnel))
	for i, t := range tunnelsRes.Payload.Tunnel {
		tunnelsList[i] = listResultTunnel{
			TunnelID:       t.TunnelID,
			ListenAgentID:  t.ListenAgentID,
			ListenPort:     uint16(t.ListenPort),
			ConnectAgentID: t.ConnectAgentID,
			ConnectPort:    uint16(t.ConnectPort),
		}
	}

	return &listTunnelsResult{
		Tunnels: tunnelsList,
	}, nil
}

// register command
var (
	ListTunnels  = new(listTunnelsCommand)
	ListTunnelsC = inventoryListC.Command("tunnels", "Show tunnels in inventory").Hide(hide)
)

func init() {
	ListTunnelsC.Flag("pmm-agent-id", "Filter by pmm-agent identifier").StringVar(&ListTunnels.filters.AgentID)
}
