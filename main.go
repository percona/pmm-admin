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

package main

import (
	"net/url"
	"os"

	"github.com/Percona-Lab/pmm-api/http/client"
	"github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/Percona-Lab/pmm-admin/commands"
)

var (
	Version = "2.0.0-dev"
)

func main() {
	app := kingpin.New("pmm-admin", "Version "+Version+".")
	app.HelpFlag.Short('h')
	pmmServerAddressF := app.Flag("server-url", "PMM Server URL.").Envar("PMM_ADMIN_SERVER_URL").Required().String()
	// debugF := app.Flag("debug", "Enable debug output.").Envar("PMM_ADMIN_DEBUG").Bool()
	kingpin.MustParse(app.Parse(os.Args[1:]))

	u, err := url.Parse(*pmmServerAddressF)
	if err != nil {
		logrus.Fatal(err)
	}
	if u.Host == "" || u.Path == "" || u.Scheme == "" {
		logrus.Fatal("Invalid PMM Server URL.")
	}
	client.Default = client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
		Host:     u.Host,
		BasePath: u.Path,
		Schemes:  []string{u.Scheme},
	})
	commonParams := commands.CommonParams{
		Client: client.Default,
	}

	cmd := commands.AddMySQLCmd{
		CommonParams: commonParams,
		Username:     "username",
		Password:     "password",
	}
	cmd.Run()
}
