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
	"strings"

	httptransport "github.com/go-openapi/runtime/client"
	inventorypb "github.com/percona/pmm/api/inventorypb/json/client"
	managementpb "github.com/percona/pmm/api/managementpb/json/client"
	serverpb "github.com/percona/pmm/api/serverpb/json/client"
	"github.com/percona/pmm/version"
	"github.com/sirupsen/logrus"
	"golang.org/x/sys/unix"
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/percona/pmm-admin/agentlocal"
	"github.com/percona/pmm-admin/commands"
	"github.com/percona/pmm-admin/commands/inventory"
	"github.com/percona/pmm-admin/commands/management"
)

func main() {
	kingpin.CommandLine.Name = "pmm-admin"
	kingpin.CommandLine.Help = fmt.Sprintf("Version %s.", version.Version)
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.CommandLine.Version(version.FullInfo())

	serverURLF := kingpin.Flag("server-url", "PMM Server URL.").String()
	serverInsecureTLSF := kingpin.Flag("server-insecure-tls", "").Bool()
	debugF := kingpin.Flag("debug", "Enable debug logging.").Bool()
	traceF := kingpin.Flag("trace", "Enable trace logging (implies debug).").Bool()
	jsonF := kingpin.Flag("json", "Enable JSON output.").Bool()

	cmd := kingpin.Parse()

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
			msg := []string{
				"Failed to get PMM Server parameters from local pmm-agent.",
				err.Error(),
				"Please use --server-url flag to specify PMM Server URL.",
			}
			logrus.Fatal(strings.Join(msg, "\n"))
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
		if httpTransport.TLSClientConfig == nil {
			httpTransport.TLSClientConfig = new(tls.Config)
		}
		httpTransport.TLSClientConfig.InsecureSkipVerify = true
	}

	inventorypb.Default.SetTransport(transport)
	managementpb.Default.SetTransport(transport)
	serverpb.Default.SetTransport(transport)

	var command commands.Command
	switch cmd {
	case management.RegisterC.FullCommand():
		command = management.Register

	case management.AddMySQLC.FullCommand():
		command = management.AddMySQL

	case inventory.AddNodeGenericC.FullCommand():
		command = inventory.AddNodeGeneric

	case inventory.AddNodeContainerC.FullCommand():
		command = inventory.AddNodeContainer

	case inventory.RemoveNodeC.FullCommand():
		command = inventory.RemoveNode

	case inventory.AddServiceMySQLC.FullCommand():
		command = inventory.AddServiceMySQL

	case inventory.AddServiceMongoDBC.FullCommand():
		command = inventory.AddServiceMongoDB

	case inventory.AddServicePostgreSQLC.FullCommand():
		command = inventory.AddServicePostgreSQL

	case inventory.RemoveServiceC.FullCommand():
		command = inventory.RemoveService

	case commands.ListC.FullCommand():
		command = commands.List

	case commands.StatusC.FullCommand():
		command = commands.Status

	default:
		logrus.Panicf("Unhandled command %q.", cmd)
	}

	res, err := command.Run()
	logrus.Debugf("Response: %#v", res)
	logrus.Debugf("Error: %#v", err)

	if err == nil {
		if *jsonF {
			b, err := json.Marshal(res)
			if err != nil {
				logrus.Fatalf("Failed to marshal result to JSON: %s.", err)
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
				logrus.Fatalf("Failed to marshal result to JSON: %s.", err)
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
