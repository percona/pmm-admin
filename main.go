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
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"os/signal"

	httptransport "github.com/go-openapi/runtime/client"
	inventory "github.com/percona/pmm/api/inventory/json/client"
	management "github.com/percona/pmm/api/managementpb/json/client"
	"github.com/percona/pmm/api/managementpb/json/client/node"
	server "github.com/percona/pmm/api/serverpb/json/client"
	"github.com/percona/pmm/version"
	"github.com/sirupsen/logrus"
	"golang.org/x/sys/unix"
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/percona/pmm-admin/agentlocal"
	"github.com/percona/pmm-admin/commands"
)

var (
	nodeTypes = map[string]string{
		"generic": node.RegisterBodyNodeTypeGENERICNODE,
	}
	nodeTypeKeys = []string{"generic"}
)

func main() {
	hostname, _ := os.Hostname()

	app := kingpin.New("pmm-admin", fmt.Sprintf("Version %s.", version.Version))
	app.HelpFlag.Short('h')
	app.Version(version.FullInfo())
	serverURLF := app.Flag("server-url", "PMM Server URL.").String()
	serverInsecureTLSF := app.Flag("server-insecure-tls", "").Bool()
	debugF := app.Flag("debug", "Enable debug logging.").Bool()
	traceF := app.Flag("trace", "Enable trace logging (implies debug).").Bool()
	jsonF := app.Flag("json", "Enable JSON output.").Bool()

	inventoryC := app.Command("inventory", "Inventory subcommands.")
	invAddC := inventoryC.Command("add", "Add subcommands.")
	_ = invAddC.Command("node", "Add Node.")
	_ = invAddC.Command("service", "Add Service.")
	_ = invAddC.Command("agent", "Add Agent.")
	invRemoveC := inventoryC.Command("remove", "Remove subcommands.")
	_ = invRemoveC.Command("node", "Remove Node.")
	_ = invRemoveC.Command("service", "Remove Service.")
	_ = invRemoveC.Command("agent", "Remove Agent.")

	_ = app.Command("status", "Show PMM Server and local pmm-agent status.")

	registerC := app.Command("register", "Register current Node at PMM Server.")
	registerNodeTypeF := registerC.Flag("node-type", "Node type.").Default(nodeTypeKeys[0]).Enum(nodeTypeKeys...)
	registerNodeNameF := registerC.Flag("node-name", "Node name.").Default(hostname).String()

	cmd := kingpin.MustParse(app.Parse(os.Args[1:]))

	logrus.SetFormatter(&logrus.TextFormatter{
		DisableTimestamp: true,
	})
	if *debugF {
		logrus.SetLevel(logrus.DebugLevel)
	}
	if *traceF {
		logrus.SetLevel(logrus.TraceLevel)
		logrus.SetReportCaller(true)
	}

	ctx, cancel := context.WithCancel(context.Background())

	// handle termination signals
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, unix.SIGTERM, unix.SIGINT)
	go func() {
		s := <-signals
		signal.Stop(signals)
		logrus.Warnf("Got %s, shutting down...", unix.SignalName(s.(unix.Signal)))
		cancel()
	}()

	agentlocal.SetTransport(ctx, *debugF || *traceF)

	var serverURL *url.URL
	var serverInsecureTLS bool
	if *serverURLF == "" {
		status, err := agentlocal.GetStatus()
		if err != nil {
			logrus.Fatal(err)
		}
		serverURL = status.ServerURL
		serverInsecureTLS = status.ServerInsecureTLS
	} else {
		var err error
		serverURL, err = url.Parse(*serverURLF)
		if err != nil {
			logrus.Fatal(err)
		}
		if serverURL.Path == "" {
			serverURL.Path = "/"
		}
		if serverURL.Host == "" || serverURL.Scheme == "" {
			logrus.Fatal("Invalid PMM Server URL.")
		}
		serverInsecureTLS = *serverInsecureTLSF
	}

	// use JSON APIs over HTTP/1.1
	transport := httptransport.New(serverURL.Host, serverURL.Path, []string{serverURL.Scheme})
	transport.SetLogger(logrus.WithField("component", "server-transport"))
	transport.SetDebug(*debugF || *traceF)
	transport.Context = ctx
	httpTransport := transport.Transport.(*http.Transport)
	httpTransport.TLSNextProto = map[string]func(string, *tls.Conn) http.RoundTripper{} // disable HTTP/2
	if serverInsecureTLS {
		httpTransport.TLSClientConfig.InsecureSkipVerify = true
	}

	inventory.Default.SetTransport(transport)
	management.Default.SetTransport(transport)
	server.Default.SetTransport(transport)

	var command commands.Command
	switch cmd {
	case registerC.FullCommand():
		command = &commands.Register{
			NodeType: nodeTypes[*registerNodeTypeF],
			NodeName: *registerNodeNameF,
		}
	default:
		logrus.Fatalf("Unexpected command %q.", cmd)
	}

	res, err := command.Run()
	logrus.Debugf("%#v", res)
	logrus.Debugf("%#v", err)

	if err == nil {
		if *jsonF {
			b, err := json.Marshal(res)
			if err != nil {
				logrus.Fatal(err)
			}
			fmt.Printf("%s\n", b)
		} else {
			fmt.Println(res.String())
		}

		return
	}

	switch err := err.(type) {
	case commands.ErrorResponse:
		e := commands.GetError(err)

		if *jsonF {
			b, err := json.Marshal(e)
			if err != nil {
				logrus.Fatal(err)
			}
			fmt.Printf("%s\n", b)
		} else {
			fmt.Println(e.Error)
		}

		os.Exit(1)

	default:
		fmt.Println(err)
		os.Exit(1)
	}
}
