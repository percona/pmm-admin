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
	"archive/zip"
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
	"time"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/percona/pmm-admin/agentlocal"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var summaryResultT = ParseTemplate(`
Agent ID: {{ .PMMAgentStatus.AgentID }}
Node ID : {{ .PMMAgentStatus.NodeID }}

PMM Server:
	URL    : {{ .PMMAgentStatus.ServerURL }}
	Version: {{ .PMMAgentStatus.ServerVersion }}

PMM-agent:
	Connected : {{ .PMMAgentStatus.Connected }}{{ if .PMMAgentStatus.Connected }}
	Time drift: {{ .PMMAgentStatus.ServerClockDrift }}
	Latency   : {{ .PMMAgentStatus.ServerLatency }}
{{ end }}
Agents:
{{ range .PMMAgentStatus.Agents }}	{{ .AgentID }} {{ .AgentType }} {{ .Status }}
{{ end }}
`)

type summaryResult struct {
	PMMAgentStatus *agentlocal.Status `json:"pmm_agent_status"`
}

func (res *summaryResult) Result() {}

func (res *summaryResult) String() string {
	return RenderTemplate(summaryResultT, res)
}

type summaryCommand struct {
	Archive         bool
	ArchiveFilename string
}

func getServerLogs(serverURL *url.URL, serverInsecureTLS bool) (*bytes.Reader, error) {
	transport := new(http.Transport)
	if serverInsecureTLS {
		transport.TLSClientConfig = &tls.Config{
			InsecureSkipVerify: true, //nolint:gosec
		}
	}
	client := &http.Client{
		Transport: transport,
	}
	u := serverURL.ResolveReference(&url.URL{
		Path: "logs.zip",
	})
	resp, err := client.Get(u.String())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer resp.Body.Close() //nolint:errcheck

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return bytes.NewReader(b), nil
}

func addServerLogs(r *bytes.Reader, zipW *zip.Writer) error {
	zipR, err := zip.NewReader(r, r.Size())
	if err != nil {
		return errors.WithStack(err)
	}
	// zipR.
}

func (cmd *summaryCommand) makeArchive(status *agentlocal.Status) (err error) {
	var f *os.File
	if f, err = os.Create(cmd.ArchiveFilename); err != nil {
		err = errors.WithStack(err)
		return
	}
	defer func() {
		if e := f.Close(); e != nil && err == nil {
			err = errors.WithStack(e)
		}
	}()

	w := zip.NewWriter(f)
	defer func() {
		if e := w.Close(); e != nil && err == nil {
			err = errors.WithStack(e)
		}
	}()

	serverZip, e := getServerLogs(status.ServerURL, status.ServerInsecureTLS)
	if e != nil {
		logrus.Debugf("Failed to get logs.zip from server: %+v", e)
	}
	if serverZip != nil {
		for _, sf := range serverZip.File {
			h, e := w.CreateHeader(&zip.FileHeader{
				Name:     path.Join("server", sf.Name),
				Method:   zip.Deflate,
				Modified: sf.Modified,
			})
			if e != nil {
				logrus.Debugf("%s", e)
				continue
			}
			sf.
		}
	}
	return nil
}

func (cmd *summaryCommand) Run() (Result, error) {
	status, err := agentlocal.GetStatus(agentlocal.RequestNetworkInfo)
	if err != nil {
		return nil, err
	}

	res := &summaryResult{
		PMMAgentStatus: status,
	}
	if !cmd.Archive {
		return res, nil
	}

	return res, nil
}

// register command
var (
	Summary  = new(summaryCommand)
	SummaryC = kingpin.Command("summary", "Show summary status information")
)

func init() {
	SummaryC.Flag("archive", "Generate summary archive file").BoolVar(&Summary.Archive)

	hostname, _ := os.Hostname()
	archiveFilename := fmt.Sprintf("summary_%s_%s.zip",
		strings.Replace(hostname, ".", "_", -1), time.Now().Format("2006_01_02_15_04_05"))
	SummaryC.Flag("archive-file", "Summary archive filename").Default(archiveFilename).StringVar(&Summary.ArchiveFilename)
}
