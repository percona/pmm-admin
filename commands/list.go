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
	"strings"

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
	Agents   []listResultAgent   `json:"agent"`
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

	var servicesList []listResultService
	for _, s := range servicesRes.Payload.Mysql {
		servicesList = append(servicesList, listResultService{
			ServiceType: services.ListServicesBodyServiceTypeMYSQLSERVICE,
			ServiceID:   s.ServiceID,
			ServiceName: s.ServiceName,
			AddressPort: net.JoinHostPort(s.Address, strconv.FormatInt(s.Port, 10)),
		})
	}
	for _, s := range servicesRes.Payload.Mongodb {
		servicesList = append(servicesList, listResultService{
			ServiceType: services.ListServicesBodyServiceTypeMONGODBSERVICE,
			ServiceID:   s.ServiceID,
			ServiceName: s.ServiceName,
			AddressPort: net.JoinHostPort(s.Address, strconv.FormatInt(s.Port, 10)),
		})
	}
	for _, s := range servicesRes.Payload.Postgresql {
		servicesList = append(servicesList, listResultService{
			ServiceType: services.ListServicesBodyServiceTypePOSTGRESQLSERVICE,
			ServiceID:   s.ServiceID,
			ServiceName: s.ServiceName,
			AddressPort: net.JoinHostPort(s.Address, strconv.FormatInt(s.Port, 10)),
		})
	}
	for _, s := range servicesRes.Payload.Proxysql {
		servicesList = append(servicesList, listResultService{
			ServiceType: services.ListServicesBodyServiceTypePROXYSQLSERVICE,
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
		res := strings.ToLower(pointer.GetString(s))
		if res == "" {
			res = "unknown"
		}
		if disabled {
			res += " (disabled)"
		}
		return res
	}

	pmmAgentIDs := map[string]struct{}{}
	var agentsList []listResultAgent
	for _, a := range agentsRes.Payload.PMMAgent {
		if a.RunsOnNodeID == cmd.NodeID {
			pmmAgentIDs[a.AgentID] = struct{}{}

			status := "disconnected"
			if a.Connected {
				status = "connected"
			}
			agentsList = append(agentsList, listResultAgent{
				AgentType: agents.ListAgentsBodyAgentTypePMMAGENT,
				AgentID:   a.AgentID,
				Status:    status,
			})
		}
	}
	for _, a := range agentsRes.Payload.NodeExporter {
		if _, ok := pmmAgentIDs[a.PMMAgentID]; ok {
			agentsList = append(agentsList, listResultAgent{
				AgentType: agents.ListAgentsBodyAgentTypeNODEEXPORTER,
				AgentID:   a.AgentID,
				Status:    getStatus(a.Status, a.Disabled),
			})
		}
	}
	for _, a := range agentsRes.Payload.MysqldExporter {
		if _, ok := pmmAgentIDs[a.PMMAgentID]; ok {
			agentsList = append(agentsList, listResultAgent{
				AgentType: agents.ListAgentsBodyAgentTypeMYSQLDEXPORTER,
				AgentID:   a.AgentID,
				ServiceID: a.ServiceID,
				Status:    getStatus(a.Status, a.Disabled),
			})
		}
	}
	for _, a := range agentsRes.Payload.MongodbExporter {
		if _, ok := pmmAgentIDs[a.PMMAgentID]; ok {
			agentsList = append(agentsList, listResultAgent{
				AgentType: agents.ListAgentsBodyAgentTypeMONGODBEXPORTER,
				AgentID:   a.AgentID,
				ServiceID: a.ServiceID,
				Status:    getStatus(a.Status, a.Disabled),
			})
		}
	}
	for _, a := range agentsRes.Payload.PostgresExporter {
		if _, ok := pmmAgentIDs[a.PMMAgentID]; ok {
			agentsList = append(agentsList, listResultAgent{
				AgentType: agents.ListAgentsBodyAgentTypePOSTGRESEXPORTER,
				AgentID:   a.AgentID,
				ServiceID: a.ServiceID,
				Status:    getStatus(a.Status, a.Disabled),
			})
		}
	}
	for _, a := range agentsRes.Payload.ProxysqlExporter {
		if _, ok := pmmAgentIDs[a.PMMAgentID]; ok {
			agentsList = append(agentsList, listResultAgent{
				AgentType: agents.ListAgentsBodyAgentTypePROXYSQLEXPORTER,
				AgentID:   a.AgentID,
				ServiceID: a.ServiceID,
				Status:    getStatus(a.Status, a.Disabled),
			})
		}
	}
	for _, a := range agentsRes.Payload.RDSExporter {
		if _, ok := pmmAgentIDs[a.PMMAgentID]; ok {
			agentsList = append(agentsList, listResultAgent{
				AgentType: agents.ListAgentsBodyAgentTypeRDSEXPORTER,
				AgentID:   a.AgentID,
				Status:    getStatus(a.Status, a.Disabled),
			})
		}
	}
	for _, a := range agentsRes.Payload.QANMysqlPerfschemaAgent {
		if _, ok := pmmAgentIDs[a.PMMAgentID]; ok {
			agentsList = append(agentsList, listResultAgent{
				AgentType: agents.ListAgentsBodyAgentTypeQANMYSQLPERFSCHEMAAGENT,
				AgentID:   a.AgentID,
				ServiceID: a.ServiceID,
				Status:    getStatus(a.Status, a.Disabled),
			})
		}
	}
	for _, a := range agentsRes.Payload.QANMysqlSlowlogAgent {
		if _, ok := pmmAgentIDs[a.PMMAgentID]; ok {
			agentsList = append(agentsList, listResultAgent{
				AgentType: agents.ListAgentsBodyAgentTypeQANMYSQLSLOWLOGAGENT,
				AgentID:   a.AgentID,
				ServiceID: a.ServiceID,
				Status:    getStatus(a.Status, a.Disabled),
			})
		}
	}
	for _, a := range agentsRes.Payload.QANMongodbProfilerAgent {
		if _, ok := pmmAgentIDs[a.PMMAgentID]; ok {
			agentsList = append(agentsList, listResultAgent{
				AgentType: agents.ListAgentsBodyAgentTypeQANMONGODBPROFILERAGENT,
				AgentID:   a.AgentID,
				ServiceID: a.ServiceID,
				Status:    getStatus(a.Status, a.Disabled),
			})
		}
	}
	for _, a := range agentsRes.Payload.QANPostgresqlPgstatementsAgent {
		if _, ok := pmmAgentIDs[a.PMMAgentID]; ok {
			agentsList = append(agentsList, listResultAgent{
				AgentType: agents.ListAgentsBodyAgentTypeQANPOSTGRESQLPGSTATEMENTSAGENT,
				AgentID:   a.AgentID,
				ServiceID: a.ServiceID,
				Status:    getStatus(a.Status, a.Disabled),
			})
		}
	}

	return &listResult{
		Services: servicesList,
		Agents:   agentsList,
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
