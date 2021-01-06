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

package management

import (
	"fmt"
	"net"
	"net/url"
	"strconv"

	"github.com/AlekSi/pointer"
	"github.com/percona/pmm/api/managementpb/json/client"
	"github.com/percona/pmm/api/managementpb/json/client/external"

	"github.com/percona/pmm-admin/commands"
)

var addExternalServerlessResultT = commands.ParseTemplate(`
External Service added.
Service ID  : {{ .Service.ServiceID }}
Service name: {{ .Service.ServiceName }}
Group       : {{ .Service.Group }}
`)

type addExternalServerlessResult struct {
	Service *external.AddExternalOKBodyService `json:"service"`
}

func (res *addExternalServerlessResult) Result() {}

func (res *addExternalServerlessResult) String() string {
	return commands.RenderTemplate(addExternalServerlessResultT, res)
}

type addExternalServerlessCommand struct {
	Name     string
	Username string
	Password string

	URL         string
	Scheme      string
	Address     string
	Host        string
	ListenPort  uint16
	MetricsPath string

	Environment    string
	Cluster        string
	ReplicationSet string
	CustomLabels   string
	Group          string

	MachineID     string
	Distro        string
	ContainerID   string
	ContainerName string
	NodeModel     string
	Region        string
	Az            string
}

func (cmd *addExternalServerlessCommand) Run() (commands.Result, error) {
	customLabels, err := commands.ParseCustomLabels(cmd.CustomLabels)
	if err != nil {
		return nil, err
	}

	scheme, metricsPath, address, port, err := cmd.processURLFlags()
	if err != nil {
		return nil, err
	}

	params := &external.AddExternalParams{
		Body: external.AddExternalBody{
			AddNode: &external.AddExternalParamsBodyAddNode{
				NodeType:      pointer.ToString("REMOTE_NODE"),
				NodeName:      cmd.Name,
				MachineID:     cmd.MachineID,
				Distro:        cmd.Distro,
				ContainerID:   cmd.ContainerID,
				ContainerName: cmd.ContainerName,
				NodeModel:     cmd.NodeModel,
				Region:        cmd.Region,
				Az:            cmd.Az,
				CustomLabels:  customLabels,
			},
			Address:        address,
			ServiceName:    cmd.Name,
			Username:       cmd.Username,
			Password:       cmd.Password,
			Scheme:         scheme,
			MetricsPath:    metricsPath,
			ListenPort:     int64(port),
			Environment:    cmd.Environment,
			Cluster:        cmd.Cluster,
			ReplicationSet: cmd.ReplicationSet,
			CustomLabels:   customLabels,
			MetricsMode:    pointer.ToString("PULL"),
			Group:          cmd.Group,
		},
		Context: commands.Ctx,
	}
	resp, err := client.Default.External.AddExternal(params)
	if err != nil {
		return nil, err
	}

	return &addExternalServerlessResult{
		Service: resp.Payload.Service,
	}, nil
}

func (cmd *addExternalServerlessCommand) processURLFlags() (scheme, metricsPath, address string, port uint16, err error) {
	scheme = cmd.Scheme
	address = cmd.Host
	port = cmd.ListenPort
	metricsPath = cmd.MetricsPath

	switch {
	case cmd.URL != "":
		uri, err := url.Parse(cmd.URL)
		if err != nil {
			return "", "", "", 0, fmt.Errorf("couldn't parse URL %s : %s", cmd.URL, err)
		}
		scheme = uri.Scheme
		address = uri.Hostname()
		portS := uri.Port()
		if portS != "" {
			portI, err := strconv.Atoi(portS)
			if err != nil {
				return "", "", "", 0, err
			}
			port = uint16(portI)
		}
		metricsPath = uri.Path
	case cmd.Address != "":
		host, portS, err := net.SplitHostPort(cmd.Address)
		if err != nil {
			return "", "", "", 0, err
		}
		address = host
		portI, err := strconv.Atoi(portS)
		if err != nil {
			return "", "", "", 0, err
		}
		port = uint16(portI)
	}

	return scheme, metricsPath, address, port, nil
}

// register command
var (
	AddExternalServerless  = new(addExternalServerlessCommand)
	AddExternalServerlessC = AddC.Command("external", "Add External to monitoring")
)

func init() {
	AddExternalServerlessC.Flag("external-name", "Service name").Default("external-serverless").StringVar(&AddExternalServerless.Name)

	AddExternalServerlessC.Flag("username", "External username").StringVar(&AddExternalServerless.Username)
	AddExternalServerlessC.Flag("password", "External password").StringVar(&AddExternalServerless.Password)

	AddExternalServerlessC.Flag("scheme", "Scheme to generate URL to exporter metrics endpoints").StringVar(&AddExternalServerless.Scheme)
	AddExternalServerlessC.Flag("url", "Full URL to exporter metrics endpoints").StringVar(&AddExternalServerless.URL)
	AddExternalServerlessC.Flag("address", "External exporter address and port").StringVar(&AddExternalServerless.Address)
	AddExternalServerlessC.Flag("host", "External exporters hostname or IP address").StringVar(&AddExternalServerless.Host)
	AddExternalServerlessC.Flag("listen-port", "Listen port of external exporter for scraping metrics.").Uint16Var(&AddExternalServerless.ListenPort)
	AddExternalServerlessC.Flag("metrics-path", "Path under which metrics are exposed, used to generate URL.").StringVar(&AddExternalServerless.MetricsPath)

	AddExternalServerlessC.Flag("environment", "Environment name").StringVar(&AddExternalServerless.Environment)
	AddExternalServerlessC.Flag("cluster", "Cluster name").StringVar(&AddExternalServerless.Cluster)
	AddExternalServerlessC.Flag("replication-set", "Replication set name").StringVar(&AddExternalServerless.ReplicationSet)
	AddExternalServerlessC.Flag("custom-labels", "Custom user-assigned labels").StringVar(&AddExternalServerless.CustomLabels)
	groupHelp := fmt.Sprintf("Group name of external service (default: %s)", defaultGroupExternalExporter)
	AddExternalServerlessC.Flag("group", groupHelp).Default(defaultGroupExternalExporter).StringVar(&AddExternalServerless.Group)

	AddExternalServerlessC.Flag("machine-id", "Node machine-id").StringVar(&AddExternalServerless.MachineID)
	AddExternalServerlessC.Flag("distro", "Node OS distribution").StringVar(&AddExternalServerless.Distro)
	AddExternalServerlessC.Flag("container-id", "Container ID").StringVar(&AddExternalServerless.ContainerID)
	AddExternalServerlessC.Flag("container-name", "Container name").StringVar(&AddExternalServerless.ContainerName)
	AddExternalServerlessC.Flag("node-model", "Node model").StringVar(&AddExternalServerless.NodeModel)
	AddExternalServerlessC.Flag("region", "Node region").StringVar(&AddExternalServerless.Region)
	AddExternalServerlessC.Flag("az", "Node availability zone").StringVar(&AddExternalServerless.Az)
}
