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

package commands

import (
	"github.com/percona/pmm/api/inventory/json/client"
	"github.com/percona/pmm/api/inventory/json/client/agents"
)

// AddMySQLCmd implements `pmm-admin add mysql` command.
type AddMySQLCmd struct {
	CommonParams
	Username string
	Password string
}

// Run implements Command interface.
func (cmd *AddMySQLCmd) Run() Result {
	// TODO get NodeID from local pmm-agent

	// TODO get or create MySQL service for this Node via pmm-managed

	params := &agents.AddMySqldExporterParams{
		Body: agents.AddMySqldExporterBody{
			// TODO RunsOnNodeID
			// TODO ServiceID
			Username: cmd.Username,
			Password: cmd.Password,
		},
	}
	resp, err := client.Default.Agents.AddMySqldExporter(params)
	if err != nil {
		cmd.Logger.Error(err)
		return nil
	}
	cmd.Logger.Infof("mysqld_exporter started on %d.", resp.Payload.MysqldExporter.ListenPort)
	return nil
}

// check interfaces
var (
	_ Cmd = (*AddMySQLCmd)(nil)
)
