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
	"github.com/sirupsen/logrus"
	"strings"
	"testing"

	"github.com/percona/pmm-admin/commands"

	mysql "github.com/percona/pmm/api/managementpb/json/client/my_sql"
	"github.com/stretchr/testify/assert"
)

func TestAddMySQL(t *testing.T) {
	t.Run("TablestatEnabled", func(t *testing.T) {
		res := &addMySQLResult{
			Service: &mysql.AddMySQLOKBodyService{
				ServiceID:   "/service_id/1",
				ServiceName: "mysql-1",
			},
			MysqldExporter: &mysql.AddMySQLOKBodyMysqldExporter{
				TablestatsGroupTableLimit: 1000,
				TablestatsGroupDisabled:   false,
			},
			TableCount: 500,
		}
		expected := strings.TrimSpace(`
MySQL Service added.
Service ID  : /service_id/1
Service name: mysql-1

Table statistics collection enabled (the limit is 1000, the actual table count is 500).
		`)
		assert.Equal(t, expected, strings.TrimSpace(res.String()))
	})

	t.Run("TablestatEnabledNoLimit", func(t *testing.T) {
		res := &addMySQLResult{
			Service: &mysql.AddMySQLOKBodyService{
				ServiceID:   "/service_id/1",
				ServiceName: "mysql-1",
			},
			MysqldExporter: &mysql.AddMySQLOKBodyMysqldExporter{
				TablestatsGroupTableLimit: 0,
				TablestatsGroupDisabled:   false,
			},
			TableCount: 2000,
		}
		expected := strings.TrimSpace(`
MySQL Service added.
Service ID  : /service_id/1
Service name: mysql-1

Table statistics collection enabled (the table count limit is not set).
		`)
		assert.Equal(t, expected, strings.TrimSpace(res.String()))
	})

	t.Run("TablestatEnabledUnknown", func(t *testing.T) {
		res := &addMySQLResult{
			Service: &mysql.AddMySQLOKBodyService{
				ServiceID:   "/service_id/1",
				ServiceName: "mysql-1",
			},
			MysqldExporter: &mysql.AddMySQLOKBodyMysqldExporter{
				TablestatsGroupTableLimit: 1000,
				TablestatsGroupDisabled:   false,
			},
			TableCount: 0,
		}
		expected := strings.TrimSpace(`
MySQL Service added.
Service ID  : /service_id/1
Service name: mysql-1

Table statistics collection enabled (the limit is 1000, the actual table count is unknown).
		`)
		assert.Equal(t, expected, strings.TrimSpace(res.String()))
	})

	t.Run("TablestatDisabled", func(t *testing.T) {
		res := &addMySQLResult{
			Service: &mysql.AddMySQLOKBodyService{
				ServiceID:   "/service_id/1",
				ServiceName: "mysql-1",
			},
			MysqldExporter: &mysql.AddMySQLOKBodyMysqldExporter{
				TablestatsGroupTableLimit: 1000,
				TablestatsGroupDisabled:   true,
				TLS:                       true,
				TLSCa:                     "ca",
				TLSCert:                   "cert",
				TLSKey:                    "key",
			},
			TableCount: 2000,
		}
		expected := strings.TrimSpace(`
MySQL Service added.
Service ID  : /service_id/1
Service name: mysql-1

Table statistics collection disabled (the limit is 1000, the actual table count is 2000).
		`)
		assert.Equal(t, expected, strings.TrimSpace(res.String()))
	})

	t.Run("TablestatDisabledAlways", func(t *testing.T) {
		res := &addMySQLResult{
			Service: &mysql.AddMySQLOKBodyService{
				ServiceID:   "/service_id/1",
				ServiceName: "mysql-1",
			},
			MysqldExporter: &mysql.AddMySQLOKBodyMysqldExporter{
				TablestatsGroupTableLimit: -1,
				TablestatsGroupDisabled:   true,
			},
			TableCount: 2000,
		}
		expected := strings.TrimSpace(`
MySQL Service added.
Service ID  : /service_id/1
Service name: mysql-1

Table statistics collection disabled (always).
		`)
		assert.Equal(t, expected, strings.TrimSpace(res.String()))
	})

	t.Run("EmptyMysqlExporter", func(t *testing.T) {
		res := &addMySQLResult{
			MysqldExporter: nil,
		}
		expected := ""
		assert.Equal(t, expected, strings.TrimSpace(res.TablestatStatus()))
	})
}

func TestRun(t *testing.T) {
	t.Run("CreateUser", func(t *testing.T) {
		cmd := &addMySQLCommand{
			CreateUser: true,
		}
		_, err := cmd.Run()

		if assert.Error(t, err) {
			expected := "Unrecognized option. To create a user, see 'https://www.percona.com/doc/percona-monitoring-and-management/2.x/concepts/services-mysql.html#pmm-conf-mysql-user-account-creating'"
			assert.Equal(t, expected, err.Error())
		}
	})
}

func TestApplyDefaults(t *testing.T) {
	t.Run("password and username is set", func(t *testing.T) {
		file, cleanup, e := commands.DefaultConfig("[client]\nuser=root\npassword=toor\n")
		if e != nil {
			t.Fatal(e)
		}
		defer cleanup()

		cmd := &addMySQLCommand{}

		commands.ConfigureDefaults(file.Name(), cmd)

		assert.Equal(t, "root", cmd.Username)
		assert.Equal(t, "toor", cmd.Password)
	})

	t.Run("password and username from config have lower priority", func(t *testing.T) {
		logrus.SetLevel(logrus.TraceLevel)
		file, cleanup, e := commands.DefaultConfig("[client]\nuser=root\npassword=toor\n")
		if e != nil {
			t.Fatal(e)
		}
		defer cleanup()

		cmd := &addMySQLCommand{
			Username: "default-username",
			Password: "default-password",
		}

		commands.ConfigureDefaults(file.Name(), cmd)

		assert.Equal(t, "default-username", cmd.Username)
		assert.Equal(t, "default-password", cmd.Password)
	})

	t.Run("not updated if not set", func(t *testing.T) {
		file, cleanup, e := commands.DefaultConfig("")
		if e != nil {
			t.Fatal(e)
		}
		defer cleanup()

		cmd := &addMySQLCommand{
			Username: "default-username",
			Password: "default-password",
		}

		commands.ConfigureDefaults(file.Name(), cmd)

		assert.Equal(t, "default-username", cmd.Username)
		assert.Equal(t, "default-password", cmd.Password)
	})

	t.Run("only username is set", func(t *testing.T) {
		file, cleanup, e := commands.DefaultConfig("[client]\nuser=root\n")
		if e != nil {
			t.Fatal(e)
		}
		defer cleanup()

		cmd := &addMySQLCommand{}

		commands.ConfigureDefaults(file.Name(), cmd)

		assert.Equal(t, "root", cmd.Username)
		assert.Equal(t, "", cmd.Password)
	})

	t.Run("only password is set", func(t *testing.T) {
		file, cleanup, e := commands.DefaultConfig("[client]\npassword=toor\n")
		if e != nil {
			t.Fatal(e)
		}
		defer cleanup()

		cmd := &addMySQLCommand{}

		commands.ConfigureDefaults(file.Name(), cmd)

		assert.Equal(t, "", cmd.Username)
		assert.Equal(t, "toor", cmd.Password)
	})
}
