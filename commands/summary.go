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
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
	"time"

	agent_local "github.com/percona/pmm/api/agentlocalpb/json/client/agent_local"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/percona/pmm-admin/agentlocal"
)

var summaryResultT = ParseTemplate(`
{{ .Filename }} created.
`)

type summaryResult struct {
	Filename string `json:"filename"`
}

func (res *summaryResult) Result() {}

func (res *summaryResult) String() string {
	return RenderTemplate(summaryResultT, res)
}

type summaryCommand struct {
	Filename string
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

func addServerData(serverURL *url.URL, serverInsecureTLS bool, zipW *zip.Writer) error {
	bytesR, err := getServerLogs(serverURL, serverInsecureTLS)
	if err != nil {
		return err
	}

	zipR, err := zip.NewReader(bytesR, bytesR.Size())
	if err != nil {
		return errors.WithStack(err)
	}

	for _, rf := range zipR.File {
		w, err := zipW.CreateHeader(&zip.FileHeader{
			Name:     path.Join("server", rf.Name),
			Method:   zip.Deflate,
			Modified: rf.Modified,
		})
		if err != nil {
			logrus.Debugf("%s", err)
			continue
		}

		r, err := rf.Open()
		if err != nil {
			logrus.Debugf("%s", err)
			continue
		}
		_, err = io.Copy(w, r)
		_ = r.Close()
		if err != nil {
			logrus.Debugf("%s", err)
			continue
		}
	}

	return nil
}

func addClientData(status *agent_local.StatusOKBody, zipW *zip.Writer) error {
	b, err := json.MarshalIndent(status, "", "  ")
	if err != nil {
		logrus.Debugf("%s", err)
	}
	w, err := zipW.CreateHeader(&zip.FileHeader{
		Name:     "client/status.json",
		Method:   zip.Deflate,
		Modified: time.Now(),
	})
	if err == nil {
		_, err = w.Write(b)
	}
	if err != nil {
		logrus.Debugf("%s", err)
	}

	return nil
}

func (cmd *summaryCommand) makeArchive() (err error) {
	var status *agent_local.StatusOKBody
	if status, err = agentlocal.GetRawStatus(context.TODO(), agentlocal.RequestNetworkInfo); err != nil {
		err = errors.WithStack(err)
		return
	}

	var f *os.File
	if f, err = os.Create(cmd.Filename); err != nil {
		err = errors.WithStack(err)
		return
	}
	defer func() {
		if e := f.Close(); e != nil && err == nil {
			err = errors.WithStack(e)
		}
	}()

	zipW := zip.NewWriter(f)
	defer func() {
		if e := zipW.Close(); e != nil && err == nil {
			err = errors.WithStack(e)
		}
	}()

	if e := addClientData(status, zipW); e != nil {
		logrus.Warnf("Failed to add client data: %s", e)
		logrus.Debugf("%+v", e)
	}

	if si := status.ServerInfo; si != nil {
		if u, e := url.Parse(si.URL); e == nil {
			if e := addServerData(u, si.InsecureTLS, zipW); e != nil {
				logrus.Warnf("Failed to add server data: %s", e)
				logrus.Debugf("%+v", e)
			}
		}
	}

	return //nolint:nakedret
}

func (cmd *summaryCommand) Run() (Result, error) {
	if err := cmd.makeArchive(); err != nil {
		return nil, err
	}

	return &summaryResult{
		Filename: cmd.Filename,
	}, nil
}

// register command
var (
	Summary  = new(summaryCommand)
	SummaryC = kingpin.Command("summary", "Fetch system data for diagnostics")
)

func init() {
	hostname, _ := os.Hostname()
	filename := fmt.Sprintf("summary_%s_%s.zip",
		strings.Replace(hostname, ".", "_", -1), time.Now().Format("2006_01_02_15_04_05"))
	SummaryC.Flag("filename", "Summary archive filename").Default(filename).StringVar(&Summary.Filename)
}
