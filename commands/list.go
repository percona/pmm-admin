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
	"net"
	"strconv"

	"github.com/AlekSi/pointer"
	"github.com/percona/pmm/api/inventorypb/json/client"
	"github.com/percona/pmm/api/inventorypb/json/client/agents"
	"github.com/percona/pmm/api/inventorypb/json/client/services"
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/percona/pmm-admin/agentlocal"
)

var listResultT = ParseTemplate(`
Service type  Service name         Address and port  Service ID
{{ range .Services }}
{{- printf "%-13s" .ServiceType }} {{ printf "%-20s" .ServiceName }} {{ printf "%-17s" .AddressPort }} {{ .ServiceID }}
{{ end }}
Agent type                  Status     Agent ID                                        Service ID
{{ range .Agents }}
{{- printf "%-27s" .AgentType }} {{ printf "%-10s" .Status }} {{ .AgentID }}  {{ .ServiceID }}
{{ end }}
`)

type listResultAgent struct {
	AgentType string `json:"agent_type"`
	AgentID   string `json:"agent_id"`
	ServiceID string `json:"service_id"`
	Status    string `json:"status"`
}

type listResultService struct {
	ServiceType string `json:"service_type"`
	ServiceID   string `json:"service_id"`
	ServiceName string `json:"service_name"`
	AddressPort string `json:"address_port"`
}

type listResult struct {
	Services []listResultService `json:"service"`
	Agents   []listResultAgent   `json:"agents"`
}

func (res *listResult) Result() {}

func (res *listResult) String() string {
	return RenderTemplate(listResultT, res)
}

type listCommand struct {
	NodeID string
}

func (cmd *listCommand) Run() (Result, error) {
	if cmd.NodeID == "" {
		status, err := agentlocal.GetStatus(agentlocal.DoNotRequestNetworkInfo)
		if err != nil {
			return nil, err
		}
		cmd.NodeID = status.NodeID
	}

	servicesRes, err := client.Default.Services.ListServices(&services.ListServicesParams{
		Body: services.ListServicesBody{
			NodeID: cmd.NodeID,
		},
		Context: Ctx,
	})
	if err != nil {
		return nil, err
	}

	var services []listResultService
	for _, s := range servicesRes.Payload.Mysql {
		services = append(services, listResultService{
			ServiceType: "MYSQL_SERVICE", // inventorypb.ServiceType_MYSQL_SERVICE.String(),
			ServiceID:   s.ServiceID,
			ServiceName: s.ServiceName,
			AddressPort: net.JoinHostPort(s.Address, strconv.FormatInt(s.Port, 10)),
		})
	}
	for _, s := range servicesRes.Payload.Mongodb {
		services = append(services, listResultService{
			ServiceType: "MONGODB_SERVICE", // inventorypb.ServiceType_MONGODB_SERVICE.String(),
			ServiceID:   s.ServiceID,
			ServiceName: s.ServiceName,
			AddressPort: net.JoinHostPort(s.Address, strconv.FormatInt(s.Port, 10)),
		})
	}
	for _, s := range servicesRes.Payload.Postgresql {
		services = append(services, listResultService{
			ServiceType: "POSTGRESQL_SERVICE", // inventorypb.ServiceType_POSTGRESQL_SERVICE.String(),
			ServiceID:   s.ServiceID,
			ServiceName: s.ServiceName,
			AddressPort: net.JoinHostPort(s.Address, strconv.FormatInt(s.Port, 10)),
		})
	}
	for _, s := range servicesRes.Payload.Proxysql {
		services = append(services, listResultService{
			ServiceType: "PROXYSQL_SERVICE", //inventorypb.ServiceType_PROXYSQL_SERVICE.String(),
			ServiceID:   s.ServiceID,
			ServiceName: s.ServiceName,
			AddressPort: net.JoinHostPort(s.Address, strconv.FormatInt(s.Port, 10)),
		})
	}

	agentsRes, err := client.Default.Agents.ListAgents(&agents.ListAgentsParams{
		Context: Ctx,
	})
	if err != nil {
		return nil, err
	}

	getStatus := func(s *string, disabled bool) string {
		res := pointer.GetString(s)
		if res == "" {
			res = "unknown"
		}
		if disabled {
			res += " (disabled)"
		}
		return res
	}

	pmmAgentIDs := map[string]struct{}{}
	var agents []listResultAgent
	for _, a := range agentsRes.Payload.PMMAgent {
		if a.RunsOnNodeID == cmd.NodeID {
			pmmAgentIDs[a.AgentID] = struct{}{}

			status := "disconnected"
			if a.Connected {
				status = "connected"
			}
			agents = append(agents, listResultAgent{
				AgentType: "PMM_AGENT", // inventorypb.AgentType_PMM_AGENT.String()
				AgentID:   a.AgentID,
				Status:    status,
			})
		}
	}
	for _, a := range agentsRes.Payload.NodeExporter {
		if _, ok := pmmAgentIDs[a.PMMAgentID]; ok {
			agents = append(agents, listResultAgent{
				AgentType: "NODE_EXPORTER", // inventorypb.AgentType_NODE_EXPORTER.String()
				AgentID:   a.AgentID,
				Status:    getStatus(a.Status, a.Disabled),
			})
		}
	}
	for _, a := range agentsRes.Payload.MysqldExporter {
		if _, ok := pmmAgentIDs[a.PMMAgentID]; ok {
			agents = append(agents, listResultAgent{
				AgentType: "MYSQLD_EXPORTER", // inventorypb.AgentType_MYSQLD_EXPORTER.String()
				AgentID:   a.AgentID,
				ServiceID: a.ServiceID,
				Status:    getStatus(a.Status, a.Disabled),
			})
		}
	}
	for _, a := range agentsRes.Payload.MongodbExporter {
		if _, ok := pmmAgentIDs[a.PMMAgentID]; ok {
			agents = append(agents, listResultAgent{
				AgentType: "MONGODB_EXPORTER", // inventorypb.AgentType_MONGODB_EXPORTER.String()
				AgentID:   a.AgentID,
				ServiceID: a.ServiceID,
				Status:    getStatus(a.Status, a.Disabled),
			})
		}
	}
	for _, a := range agentsRes.Payload.PostgresExporter {
		if _, ok := pmmAgentIDs[a.PMMAgentID]; ok {
			agents = append(agents, listResultAgent{
				AgentType: "POSTGRES_EXPORTER", // inventorypb.AgentType_POSTGRES_EXPORTER.String()
				AgentID:   a.AgentID,
				ServiceID: a.ServiceID,
				Status:    getStatus(a.Status, a.Disabled),
			})
		}
	}
	for _, a := range agentsRes.Payload.ProxysqlExporter {
		if _, ok := pmmAgentIDs[a.PMMAgentID]; ok {
			agents = append(agents, listResultAgent{
				AgentType: "PROXYSQL_EXPORTER", // inventorypb.AgentType_PROXYSQL_EXPORTER.String()
				AgentID:   a.AgentID,
				ServiceID: a.ServiceID,
				Status:    getStatus(a.Status, a.Disabled),
			})
		}
	}
	for _, a := range agentsRes.Payload.RDSExporter {
		if _, ok := pmmAgentIDs[a.PMMAgentID]; ok {
			agents = append(agents, listResultAgent{
				AgentType: "RDS_EXPORTER", // inventorypb.AgentType_RDS_EXPORTER.String()
				AgentID:   a.AgentID,
				Status:    getStatus(a.Status, a.Disabled),
			})
		}
	}
	for _, a := range agentsRes.Payload.QANMysqlPerfschemaAgent {
		if _, ok := pmmAgentIDs[a.PMMAgentID]; ok {
			agents = append(agents, listResultAgent{
				AgentType: "QAN_MYSQL_PERFSCHEMA_AGENT", // inventorypb.AgentType_QAN_MYSQL_PERFSCHEMA_AGENT.String()
				AgentID:   a.AgentID,
				ServiceID: a.ServiceID,
				Status:    getStatus(a.Status, a.Disabled),
			})
		}
	}
	for _, a := range agentsRes.Payload.QANMysqlSlowlogAgent {
		if _, ok := pmmAgentIDs[a.PMMAgentID]; ok {
			agents = append(agents, listResultAgent{
				AgentType: "QAN_MYSQL_SLOWLOG_AGENT", // inventorypb.AgentType_QAN_MYSQL_SLOWLOG_AGENT.String()
				AgentID:   a.AgentID,
				ServiceID: a.ServiceID,
				Status:    getStatus(a.Status, a.Disabled),
			})
		}
	}
	for _, a := range agentsRes.Payload.QANMongodbProfilerAgent {
		if _, ok := pmmAgentIDs[a.PMMAgentID]; ok {
			agents = append(agents, listResultAgent{
				AgentType: "QAN_MONGODB_PROFILER_AGENT", // inventorypb.AgentType_QAN_MONGODB_PROFILER_AGENT.String()
				AgentID:   a.AgentID,
				ServiceID: a.ServiceID,
				Status:    getStatus(a.Status, a.Disabled),
			})
		}
	}
	for _, a := range agentsRes.Payload.QANPostgresqlPgstatementsAgent {
		if _, ok := pmmAgentIDs[a.PMMAgentID]; ok {
			agents = append(agents, listResultAgent{
				AgentType: "QAN_POSTGRESQL_PGSTATEMENTS_AGENT", // inventorypb.AgentType_QAN_POSTGRESQL_PGSTATEMENTS_AGENT.String()
				AgentID:   a.AgentID,
				ServiceID: a.ServiceID,
				Status:    getStatus(a.Status, a.Disabled),
			})
		}
	}

	return &listResult{
		Services: services,
		Agents:   agents,
	}, nil
}

// register command
var (
	List  = new(listCommand)
	ListC = kingpin.Command("list", "Show Services and Agents running on this Node")
)

func init() {
	ListC.Flag("node-id", "Node ID (default is autodetected)").StringVar(&List.NodeID)
}
