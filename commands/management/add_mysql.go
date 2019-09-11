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

	"github.com/alecthomas/units"

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

	AddNode       bool
	AddNodeParams addNodeParams

	// TODO remove once https://jira.percona.com/browse/PMM-4255 is done
	UsePerfschema bool
	UseSlowLog    bool

	SkipConnectionCheck  bool
	MaxSlowlogFileSize   units.Base2Bytes
	DisableQueryExamples bool
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

			SkipConnectionCheck:  cmd.SkipConnectionCheck,
			MaxSlowlogFileSize:   fmt.Sprintf("%d", cmd.MaxSlowlogFileSize),
			DisableQueryExamples: cmd.DisableQueryExamples,
		},
		Context: commands.Ctx,
	}

	if cmd.NodeName != "" {
		if cmd.AddNode {
			nodeCustomLabels, err := commands.ParseCustomLabels(cmd.AddNodeParams.CustomLabels)
			if err != nil {
				return nil, err
			}
			params.Body.AddNode = &mysql.AddMySQLParamsBodyAddNode{
				Az:            cmd.AddNodeParams.Az,
				ContainerID:   cmd.AddNodeParams.ContainerID,
				ContainerName: cmd.AddNodeParams.ContainerName,
				CustomLabels:  nodeCustomLabels,
				Distro:        cmd.AddNodeParams.Distro,
				MachineID:     cmd.AddNodeParams.MachineID,
				NodeModel:     cmd.AddNodeParams.NodeModel,
				NodeName:      cmd.NodeName,
				NodeType:      pointer.ToString(nodeTypes[cmd.AddNodeParams.NodeType]),
				Region:        cmd.AddNodeParams.Region,
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

	AddMySQLC.Flag("node-id", "Node ID (default is autodetected)").StringVar(&AddMySQL.NodeID)
	AddMySQLC.Flag("pmm-agent-id", "The pmm-agent identifier which runs this instance (default is autodetected)").StringVar(&AddMySQL.PMMAgentID)

	querySources := []string{"slowlog", "perfschema"} // TODO add "auto"
	querySourceHelp := fmt.Sprintf("Source of SQL queries, one of: %s (default: %s)", strings.Join(querySources, ", "), querySources[0])
	AddMySQLC.Flag("query-source", querySourceHelp).Default(querySources[0]).EnumVar(&AddMySQL.QuerySource, querySources...)
	AddMySQLC.Flag("use-perfschema", "Run QAN perf schema agent").Hidden().BoolVar(&AddMySQL.UsePerfschema)
	AddMySQLC.Flag("use-slowlog", "Run QAN slow log agent").Hidden().BoolVar(&AddMySQL.UseSlowLog)
	AddMySQLC.Flag("size-slow-logs", "Rotate slow logs. (0 = no rotation)").Default("1GB").BytesVar(&AddMySQL.MaxSlowlogFileSize)
	AddMySQLC.Flag("disable-query-examples", "Skip query examples").BoolVar(&AddMySQL.DisableQueryExamples)

	AddMySQLC.Flag("environment", "Environment name").StringVar(&AddMySQL.Environment)
	AddMySQLC.Flag("cluster", "Cluster name").StringVar(&AddMySQL.Cluster)
	AddMySQLC.Flag("replication-set", "Replication set name").StringVar(&AddMySQL.ReplicationSet)
	AddMySQLC.Flag("custom-labels", "Custom user-assigned labels").StringVar(&AddMySQL.CustomLabels)

	AddMySQLC.Flag("skip-connection-check", "Skip connection check").BoolVar(&AddMySQL.SkipConnectionCheck)

	AddMySQLC.Flag("add-node", "Add new node").BoolVar(&AddMySQL.AddNode)
	AddMySQLC.Flag("node-name", "Node name").StringVar(&AddMySQL.NodeName)
	addNodeFlags(AddMySQLC, &AddMySQL.AddNodeParams)
}
