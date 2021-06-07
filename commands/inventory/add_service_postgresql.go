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
	"fmt"

	"github.com/percona/pmm/api/inventorypb/json/client"
	"github.com/percona/pmm/api/inventorypb/json/client/services"

	"github.com/percona/pmm-admin/commands"
)

var addServicePostgreSQLResultT = commands.ParseTemplate(`
PostgreSQL Service added.
Service ID     : {{ .Service.ServiceID }}
Service name   : {{ .Service.ServiceName }}
Node ID        : {{ .Service.NodeID }}
{{ if .Service.Socket -}}
Socket         : {{ .Service.Socket }}
{{- else -}}
Address        : {{ .Service.Address }}
Port           : {{ .Service.Port }}
{{- end }}
Environment    : {{ .Service.Environment }}
Cluster name   : {{ .Service.Cluster }}
Replication set: {{ .Service.ReplicationSet }}
Custom labels  : {{ .Service.CustomLabels }}
`)

type addServicePostgreSQLResult struct {
	Service *services.AddPostgreSQLServiceOKBodyPostgresql `json:"postgresql"`
}

func (res *addServicePostgreSQLResult) Result() {}

func (res *addServicePostgreSQLResult) String() string {
	return commands.RenderTemplate(addServicePostgreSQLResultT, res)
}

type addServicePostgreSQLCommand struct {
	ServiceName    string
	NodeID         string
	Address        string
	Port           int64
	Socket         string
	Environment    string
	Cluster        string
	ReplicationSet string
	CustomLabels   string
	// TLS parameters
	TLS           bool
	TLSSkipVerify bool
	TLSCAFile     string
	TLSCertFile   string
	TLSKeyFile    string
}

func (cmd *addServicePostgreSQLCommand) Run() (commands.Result, error) {
	customLabels, err := commands.ParseCustomLabels(cmd.CustomLabels)
	if err != nil {
		return nil, err
	}

	var tlsCa, tlsCert, tlsKey string
	if cmd.TLS {
		if cmd.TLSCAFile == "" || cmd.TLSCertFile == "" || cmd.TLSKeyFile == "" {
			return nil, fmt.Errorf("TLS is on. You must also define tls-ca, tls-cert and tls-key flags.")
		}

		tlsCa, err = commands.ReadFile(cmd.TLSCAFile)
		if err != nil {
			return nil, err
		}

		tlsCert, err = commands.ReadFile(cmd.TLSCertFile)
		if err != nil {
			return nil, err
		}

		tlsKey, err = commands.ReadFile(cmd.TLSKeyFile)
		if err != nil {
			return nil, err
		}
	}

	params := &services.AddPostgreSQLServiceParams{
		Body: services.AddPostgreSQLServiceBody{
			ServiceName:    cmd.ServiceName,
			NodeID:         cmd.NodeID,
			Address:        cmd.Address,
			Port:           cmd.Port,
			Socket:         cmd.Socket,
			Environment:    cmd.Environment,
			Cluster:        cmd.Cluster,
			ReplicationSet: cmd.ReplicationSet,
			CustomLabels:   customLabels,

			TLS:           cmd.TLS,
			TLSCa:         tlsCa,
			TLSCert:       tlsCert,
			TLSKey:        tlsKey,
			TLSSkipVerify: cmd.TLSSkipVerify,
		},
		Context: commands.Ctx,
	}

	resp, err := client.Default.Services.AddPostgreSQLService(params)
	if err != nil {
		return nil, err
	}
	return &addServicePostgreSQLResult{
		Service: resp.Payload.Postgresql,
	}, nil
}

// register command
var (
	AddServicePostgreSQL  = new(addServicePostgreSQLCommand)
	AddServicePostgreSQLC = addServiceC.Command("postgresql", "Add PostgreSQL service to inventory").Hide(hide)
)

func init() {
	AddServicePostgreSQLC.Arg("name", "Service name").StringVar(&AddServicePostgreSQL.ServiceName)
	AddServicePostgreSQLC.Arg("node-id", "Node ID").StringVar(&AddServicePostgreSQL.NodeID)
	AddServicePostgreSQLC.Arg("address", "Address").StringVar(&AddServicePostgreSQL.Address)
	AddServicePostgreSQLC.Arg("port", "Port").Int64Var(&AddServicePostgreSQL.Port)
	AddServicePostgreSQLC.Flag("socket", "Path to socket").StringVar(&AddServicePostgreSQL.Socket)

	AddServicePostgreSQLC.Flag("environment", "Environment name").StringVar(&AddServicePostgreSQL.Environment)
	AddServicePostgreSQLC.Flag("cluster", "Cluster name").StringVar(&AddServicePostgreSQL.Cluster)
	AddServicePostgreSQLC.Flag("replication-set", "Replication set name").StringVar(&AddServicePostgreSQL.ReplicationSet)
	AddServicePostgreSQLC.Flag("custom-labels", "Custom user-assigned labels").StringVar(&AddServicePostgreSQL.CustomLabels)

	AddServicePostgreSQLC.Flag("tls", "Use TLS to connect to the database").BoolVar(&AddServicePostgreSQL.TLS)
	AddServicePostgreSQLC.Flag("tls-ca-file", "TLS CA certificate file").StringVar(&AddServicePostgreSQL.TLSCAFile)
	AddServicePostgreSQLC.Flag("tls-cert-file", "TLS certificate file").StringVar(&AddServicePostgreSQL.TLSCertFile)
	AddServicePostgreSQLC.Flag("tls-key-file", "TLS certificate key file").StringVar(&AddServicePostgreSQL.TLSKeyFile)
	AddServicePostgreSQLC.Flag("tls-skip-verify", "Skip TLS certificates validation").BoolVar(&AddServicePostgreSQL.TLSSkipVerify)
}
