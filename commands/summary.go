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
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/percona/pmm/api/serverpb/json/client"
	"github.com/percona/pmm/api/serverpb/json/client/server"
	"github.com/percona/pmm/version"
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
	Filename   string
	SkipServer bool
	Pprof      bool
}

func getServerLogs(ctx context.Context) (*bytes.Reader, error) {
	var buffer bytes.Buffer
	_, err := client.Default.Server.Logs(&server.LogsParams{Context: ctx}, &buffer)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return bytes.NewReader(buffer.Bytes()), nil
}

func addServerData(ctx context.Context, zipW *zip.Writer) error {
	bytesR, err := getServerLogs(ctx)
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

func addFileToZip(zipW *zip.Writer, fpath, name string) {
	if name == "" {
		return
	}

	b, err := ioutil.ReadFile(name) //nolint:gosec
	if err != nil {
		logrus.Debugf("%s", err)
		b = []byte(err.Error())
	}
	m := time.Now()
	if fi, _ := os.Stat(name); fi != nil {
		m = fi.ModTime()
	}

	writeFileToZipWithTime(zipW, path.Join(fpath, filepath.Base(name)), m, b)
}

func addClientCommand(zipW *zip.Writer, name string, cmd Command) {
	var b []byte
	res, err := cmd.Run()
	if res != nil {
		b = append([]byte(res.String()), "\n\n"...)
	}
	if err != nil {
		b = append(b, err.Error()...)
	}

	writeFileToZip(zipW, path.Join("client", name), b)
}

func addClientData(ctx context.Context, zipW *zip.Writer) error {
	status, err := agentlocal.GetRawStatus(ctx, agentlocal.RequestNetworkInfo)
	if err != nil {
		return err
	}

	b, err := json.MarshalIndent(status, "", "  ")
	if err != nil {
		logrus.Debugf("%s", err)
		b = []byte(err.Error())
	}
	b = append(b, '\n')
	writeFileToZip(zipW, "client/status.json", b)

	// FIXME get it via pmm-agent's API - it is _not_ a good idea to use exec there
	// golangli-lint should continue complain about it until it is fixed
	b, err = exec.Command("pmm-agent", "--version").CombinedOutput()
	if err != nil {
		logrus.Debugf("%s", err)
		b = []byte(err.Error())
	}

	writeFileToZip(zipW, "client/pmm-agent-version.txt", b)
	writeFileToZip(zipW, "client/pmm-admin-version.txt", []byte(version.FullInfo()))

	addFileToZip(zipW, "client", status.ConfigFilepath)

	addClientCommand(zipW, "list.txt", &listCommand{NodeID: status.RunsOnNodeID})

	return nil
}

func writeFileToZip(zipW *zip.Writer, fileName string, data []byte) {
	modifiedTime := time.Now()
	writeFileToZipWithTime(zipW, fileName, modifiedTime, data)
}

func writeFileToZipWithTime(zipW *zip.Writer, fileName string, modifiedTime time.Time, data []byte) {
	w, err := zipW.CreateHeader(&zip.FileHeader{
		Name:     fileName,
		Method:   zip.Deflate,
		Modified: modifiedTime,
	})
	if err == nil {
		_, err = w.Write(data)
	}
	if err != nil {
		logrus.Errorf("%s", err)
	}
}

func (cmd *summaryCommand) makeArchive(ctx context.Context) (err error) {
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

	if e := addClientData(ctx, zipW); e != nil {
		logrus.Warnf("Failed to add client data: %s", e)
		logrus.Debugf("%+v", e)
	}

	if cmd.Pprof {
		files := cmd.getPprofData(ctx)
		for _, file := range files {
			writeFileToZip(zipW, path.Join("pprof", file.name), file.body)
		}
	}

	if !cmd.SkipServer {
		if e := addServerData(ctx, zipW); e != nil {
			logrus.Warnf("Failed to add server data: %s", e)
			logrus.Debugf("%+v", e)
		}
	}

	return //nolint:nakedret
}

type profilerPath struct {
	suffix  string
	webPath string
}

type pprofFile struct {
	name string
	body []byte
}

func (cmd *summaryCommand) getPprofData(ctx context.Context) []pprofFile {
	profilerPaths := []profilerPath{
		{
			suffix:  "profile.pb.gz",
			webPath: "/profile?seconds=60",
		},
		{
			suffix:  "heap.pb.gz",
			webPath: "/heap?gc=1",
		},
		{
			suffix:  "trace.out",
			webPath: "/trace?seconds=10",
		},
	}
	apps := map[string]string{
		"pmm-agent": "http://127.0.0.1:7777/debug/pprof",
	}

	if !cmd.SkipServer {
		apps["pmm-managed"] = "http://127.0.0.1:7773/debug/pprof"
		apps["qan-api2"] = "http://127.0.0.1:9933/debug/pprof"
	}

	out := make(chan pprofFile)
	var files []pprofFile
	var wg sync.WaitGroup
	var wgr sync.WaitGroup

	wgr.Add(1)

	go func() {
		for fileName := range out {
			files = append(files, fileName)
		}

		wgr.Done()
	}()

	//We download data from different apps in parallel, but don't download profiles from the same app in parallel.
	for appName, baseURL := range apps {
		wg.Add(1)

		go func(appName, baseURL string) {
			defer wg.Done()

			for _, file := range profilerPaths {
				fs := fmt.Sprintf("%s-%s", appName, file.suffix)
				url := baseURL + file.webPath

				logrus.Debugf("Started downloading profiler data from %s ...", url)
				b, err := getURL(ctx, url)
				if err != nil {
					logrus.Debugf("Can't download profiler data from %s: %s.", url, err)
					continue
				}

				logrus.Debugf("Finished downloading profiler data from %s.", url)
				out <- pprofFile{
					name: fs,
					body: b,
				}
			}
		}(appName, baseURL)
	}

	wg.Wait()
	close(out)
	wgr.Wait()

	return files
}

func getURL(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer resp.Body.Close() //nolint:errcheck

	if resp.StatusCode != http.StatusOK {
		return nil, errors.Errorf("status code: %d", resp.StatusCode)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "cannot read response body")
	}
	return b, nil
}

// TODO remove
func (cmd *summaryCommand) Run() (Result, error) {
	return cmd.RunWithContext(context.TODO())
}

func (cmd *summaryCommand) RunWithContext(ctx context.Context) (Result, error) {
	if err := cmd.makeArchive(ctx); err != nil {
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
	SummaryC.Flag("skip-server", "Skip fetching logs.zip from PMM Server").BoolVar(&Summary.SkipServer)
	SummaryC.Flag("pprof", "Include profiler information in the logs").BoolVar(&Summary.Pprof)
}
