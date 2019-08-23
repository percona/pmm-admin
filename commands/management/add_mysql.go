// pmm-admin
// Copyright (C) 2018 Percona LLC
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

package management

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/AlekSi/pointer"
	"github.com/percona/pmm/api/managementpb/json/client"
	mysql "github.com/percona/pmm/api/managementpb/json/client/my_sql"

	"github.com/percona/pmm-admin/agentlocal"
	"github.com/percona/pmm-admin/commands"
)

var addMySQLResultT = commands.ParseTemplate(`
MySQL Service added.
Service ID  : {{ .Service.ServiceID }}
Service name: {{ .Service.ServiceName }}
`)

type addMySQLResult struct {
	Service *mysql.AddMySQLOKBodyService `json:"service"`
}

func (res *addMySQLResult) Result() {}

func (res *addMySQLResult) String() string {
	return commands.RenderTemplate(addMySQLResultT, res)
}

type addMySQLCommand struct {
	AddressPort    string
	NodeID         string
	NodeName       string
	PMMAgentID     string
	ServiceName    string
	Username       string
	Password       string
	Environment    string
	Cluster        string
	ReplicationSet string
	CustomLabels   string

	QuerySource string

	Register       bool
	RegisterParams registerCommand

	// TODO remove once https://jira.percona.com/browse/PMM-4255 is done
	UsePerfschema bool
	UseSlowLog    bool

	SkipConnectionCheck bool
}

func (cmd *addMySQLCommand) Run() (commands.Result, error) {
	customLabels, err := commands.ParseCustomLabels(cmd.CustomLabels)
	if err != nil {
		return nil, err
	}

	if cmd.PMMAgentID == "" || (cmd.NodeID == "" && cmd.NodeName == "") {
		status, err := agentlocal.GetStatus(agentlocal.DoNotRequestNetworkInfo)
		if err != nil {
			return nil, err
		}
		if cmd.PMMAgentID == "" {
			cmd.PMMAgentID = status.AgentID
		}
		if cmd.NodeID == "" && cmd.NodeName == "" {
			cmd.NodeID = status.NodeID
		}
	}

	host, portS, err := net.SplitHostPort(cmd.AddressPort)
	if err != nil {
		return nil, err
	}
	port, err := strconv.Atoi(portS)
	if err != nil {
		return nil, err
	}

	// ignore query source if old flags are present for compatibility
	useSlowLog, usePerfschema := cmd.UseSlowLog, cmd.UsePerfschema
	if !(useSlowLog || usePerfschema) {
		switch cmd.QuerySource {
		case "slowlog":
			useSlowLog = true
		case "perfschema":
			usePerfschema = true
		}
	}

	params := &mysql.AddMySQLParams{
		Body: mysql.AddMySQLBody{
			NodeID:         cmd.NodeID,
			ServiceName:    cmd.ServiceName,
			Address:        host,
			Port:           int64(port),
			PMMAgentID:     cmd.PMMAgentID,
			Environment:    cmd.Environment,
			Cluster:        cmd.Cluster,
			ReplicationSet: cmd.ReplicationSet,
			Username:       cmd.Username,
			Password:       cmd.Password,
			CustomLabels:   customLabels,

			QANMysqlSlowlog:    useSlowLog,
			QANMysqlPerfschema: usePerfschema,

			SkipConnectionCheck: cmd.SkipConnectionCheck,
		},
		Context: commands.Ctx,
	}
	if cmd.NodeName != "" {
		if cmd.Register {
			nodeCustomLabels, err := commands.ParseCustomLabels(cmd.RegisterParams.CustomLabels)
			if err != nil {
				return nil, err
			}
			params.Body.RegisterNode = &mysql.AddMySQLParamsBodyRegisterNode{
				Address:       cmd.RegisterParams.Address,
				Az:            cmd.RegisterParams.Az,
				ContainerID:   cmd.RegisterParams.ContainerID,
				ContainerName: cmd.RegisterParams.ContainerName,
				CustomLabels:  nodeCustomLabels,
				Distro:        cmd.RegisterParams.Distro,
				MachineID:     cmd.RegisterParams.MachineID,
				NodeModel:     cmd.RegisterParams.NodeModel,
				NodeName:      cmd.NodeName,
				NodeType:      pointer.ToString(nodeTypes[cmd.RegisterParams.NodeType]),
				Region:        cmd.RegisterParams.Region,
			}
		} else {
			params.Body.NodeName = cmd.NodeName
		}
	}
	resp, err := client.Default.MySQL.AddMySQL(params)
	if err != nil {
		return nil, err
	}

	return &addMySQLResult{
		Service: resp.Payload.Service,
	}, nil
}

// register command
var (
	AddMySQL  = new(addMySQLCommand)
	AddMySQLC = AddC.Command("mysql", "Add MySQL to monitoring")
)

func init() {
	AddMySQLC.Arg("address", "MySQL address and port (default: 127.0.0.1:3306").Default("127.0.0.1:3306").StringVar(&AddMySQL.AddressPort)

	hostname, _ := os.Hostname()
	serviceName := hostname + "-mysql"
	serviceNameHelp := fmt.Sprintf("Service name (autodetected default: %s)", serviceName)
	AddMySQLC.Arg("name", serviceNameHelp).Default(serviceName).StringVar(&AddMySQL.ServiceName)

	AddMySQLC.Flag("username", "MySQL username").Default("root").StringVar(&AddMySQL.Username)
	AddMySQLC.Flag("password", "MySQL password").StringVar(&AddMySQL.Password)

	querySources := []string{"slowlog", "perfschema"} // TODO add "auto"
	querySourceHelp := fmt.Sprintf("Source of SQL queries, one of: %s (default: %s)", strings.Join(querySources, ", "), querySources[0])
	AddMySQLC.Flag("query-source", querySourceHelp).Default(querySources[0]).EnumVar(&AddMySQL.QuerySource, querySources...)
	AddMySQLC.Flag("use-perfschema", "Run QAN perf schema agent").Hidden().BoolVar(&AddMySQL.UsePerfschema)
	AddMySQLC.Flag("use-slowlog", "Run QAN slow log agent").Hidden().BoolVar(&AddMySQL.UseSlowLog)

	AddMySQLC.Flag("environment", "Environment name").StringVar(&AddMySQL.Environment)
	AddMySQLC.Flag("cluster", "Cluster name").StringVar(&AddMySQL.Cluster)
	AddMySQLC.Flag("replication-set", "Replication set name").StringVar(&AddMySQL.ReplicationSet)
	AddMySQLC.Flag("custom-labels", "Custom user-assigned labels").StringVar(&AddMySQL.CustomLabels)

	AddMySQLC.Flag("skip-connection-check", "Skip connection check").BoolVar(&AddMySQL.SkipConnectionCheck)

	AddMySQLC.Flag("register-node", "Register new node").BoolVar(&AddMySQL.Register)

	AddMySQLC.Arg("node-address", "Node address").StringVar(&AddMySQL.RegisterParams.Address)

	nodeTypeDefault := "remote"
	nodeTypeHelp := fmt.Sprintf("Node type, one of: %s (default: %s)", strings.Join(nodeTypeKeys, ", "), nodeTypeDefault)
	AddMySQLC.Arg("node-type", nodeTypeHelp).Default(nodeTypeDefault).EnumVar(&AddMySQL.RegisterParams.NodeType, nodeTypeKeys...)

	AddMySQLC.Flag("node-name", "Node name").StringVar(&AddMySQL.NodeName)
	AddMySQLC.Flag("machine-id", "Node machine-id (default is autodetected)").StringVar(&AddMySQL.RegisterParams.MachineID)
	AddMySQLC.Flag("distro", "Node OS distribution (default is autodetected)").StringVar(&AddMySQL.RegisterParams.Distro)
	AddMySQLC.Flag("container-id", "Container ID").StringVar(&AddMySQL.RegisterParams.ContainerID)
	AddMySQLC.Flag("container-name", "Container name").StringVar(&AddMySQL.RegisterParams.ContainerName)
	AddMySQLC.Flag("node-model", "Node model").StringVar(&AddMySQL.RegisterParams.NodeModel)
	AddMySQLC.Flag("region", "Node region").StringVar(&AddMySQL.RegisterParams.Region)
	AddMySQLC.Flag("az", "Node availability zone").StringVar(&AddMySQL.RegisterParams.Az)
	AddMySQLC.Flag("node-custom-labels", "Custom user-assigned labels").StringVar(&AddMySQL.RegisterParams.CustomLabels)
}
