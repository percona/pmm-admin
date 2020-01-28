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
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"sync"
	"time"

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

type pproofResult struct {
	filename string
	err      error
}

func (res *summaryResult) Result() {}

func (res *summaryResult) String() string {
	return RenderTemplate(summaryResultT, res)
}

type summaryCommand struct {
	Filename   string
	SkipServer bool
	Pproof     bool
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

	if resp.StatusCode != http.StatusOK {
		return nil, errors.Errorf("status code %d", resp.StatusCode)
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return bytes.NewReader(b), nil
}

func addServerData(zipW *zip.Writer, serverURL *url.URL, serverInsecureTLS bool) error {
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

	w, err := zipW.CreateHeader(&zip.FileHeader{
		Name:     path.Join(fpath, filepath.Base(name)),
		Method:   zip.Deflate,
		Modified: m,
	})
	if err == nil {
		_, err = w.Write(b)
	}
	if err != nil {
		logrus.Debugf("%s", err)
	}
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

	w, err := zipW.CreateHeader(&zip.FileHeader{
		Name:     path.Join("client", name),
		Method:   zip.Deflate,
		Modified: time.Now(),
	})
	if err == nil {
		_, err = w.Write(b)
	}
	if err != nil {
		logrus.Debugf("%s", err)
	}
}

func addClientData(zipW *zip.Writer) error {
	status, err := agentlocal.GetRawStatus(context.TODO(), agentlocal.RequestNetworkInfo)
	if err != nil {
		return err
	}

	b, err := json.MarshalIndent(status, "", "  ")
	if err != nil {
		logrus.Debugf("%s", err)
		b = []byte(err.Error())
	}
	b = append(b, '\n')
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

	w, err = zipW.CreateHeader(&zip.FileHeader{
		Name:     "client/pmm-admin-version.txt",
		Method:   zip.Deflate,
		Modified: time.Now(),
	})
	if err == nil {
		_, err = w.Write([]byte(version.FullInfo()))
	}
	if err != nil {
		logrus.Debugf("%s", err)
	}

	// FIXME get it via pmm-agent's API - it is _not_ a good idea to use exec there
	// golangli-lint should continue complain about it until it is fixed
	b, err = exec.Command("pmm-agent", "--version").CombinedOutput()
	if err != nil {
		logrus.Debugf("%s", err)
		b = []byte(err.Error())
	}
	w, err = zipW.CreateHeader(&zip.FileHeader{
		Name:     "client/pmm-agent-version.txt",
		Method:   zip.Deflate,
		Modified: time.Now(),
	})
	if err == nil {
		_, err = w.Write(b)
	}
	if err != nil {
		logrus.Debugf("%s", err)
	}

	addFileToZip(zipW, "client", status.ConfigFilepath)

	addClientCommand(zipW, "list.txt", &listCommand{NodeID: status.RunsOnNodeID})

	return nil
}

func (cmd *summaryCommand) makeArchive() (err error) {
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

	if e := addClientData(zipW); e != nil {
		logrus.Warnf("Failed to add client data: %s", e)
		logrus.Debugf("%+v", e)
	}

	if cmd.Pproof {
		files := getPprofData()
		for _, file := range files {
			addFileToZip(zipW, "pproof", file)
		}
	}

	if !cmd.SkipServer {
		if e := addServerData(zipW, GlobalFlags.ServerURL, GlobalFlags.ServerInsecureTLS); e != nil {
			logrus.Warnf("Failed to add server data: %s", e)
			logrus.Debugf("%+v", e)
		}
	}

	return //nolint:nakedret
}

func checkPproofEnabled(url string) (bool, error) {
	resp, err := http.Get(url) //nolint:G017
	if err != nil {
		return false, errors.Wrap(err, "cannot check if proflier is enabled in server & client")
	}
	defer resp.Body.Close() //nolint:errcheck

	if resp.StatusCode != http.StatusOK {
		return false, nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, errors.Wrap(err, "cannot check if proflier is enabled in server & client (error reading body")
	}

	if bytes.Contains(bytes.ToLower(body), []byte("types of profiles available")) {
		return true, nil
	}

	return false, nil
}

func getPprofData() []string {
	apps := []struct {
		app           string
		baseURL       string
		profilerPaths map[string]string
	}{
		{
			app:     "pmm-agent",
			baseURL: "http://127.0.0.1:7777/debug/pprof/",
			profilerPaths: map[string]string{
				"pmm-agent-profile.pb.gz": "/profile?seconds=60", // (profile.pb.gz)
				"pmm-agent-heap.pb.gz":    "/heap?gc=1",          //(heap.pb.gz)
				"pmm-agent-trace.out":     "/trace?seconds=10",   // (trace.out)
			},
		},
		{
			app:     "pmm-managed",
			baseURL: "http://127.0.0.1:7773/debug/pprof/",
			profilerPaths: map[string]string{
				"pmm-managed-profile.pb.gz": "/profile?seconds=60", // (profile.pb.gz)
				"pmm-managed-heap.pb.gz":    "/heap?gc=1",          //(heap.pb.gz)
				"pmm-managed-trace.out":     "/trace?seconds=10",   // (trace.out)
			},
		},
		{
			app:     "qan-api2",
			baseURL: "http://127.0.0.1:9933/debug/pprof/",
			profilerPaths: map[string]string{
				"qan-api2-profile.pb.gz": "/profile?seconds=60", // (profile.pb.gz)
				"qan-api2-heap.pb.gz":    "/heap?gc=1",          //(heap.pb.gz)
				"qan-api2-trace.out":     "/trace?seconds=10",   // (trace.out)
			},
		},
	}

	out := make(chan pproofResult)
	files := make([]string, 0)
	wg := &sync.WaitGroup{}
	wgr := &sync.WaitGroup{}

	wgr.Add(1)

	go func() {
		for result := range out {
			if result.err != nil {
				logrus.Errorf("cannot get profiles info: %s", result.err)
				continue
			}
			files = append(files, result.filename)
		}

		wgr.Done()
	}()

	for _, appInfo := range apps {
		enabled, err := checkPproofEnabled(appInfo.baseURL)
		if err != nil {
			logrus.Errorf("cannot get %s profiler status: %s", appInfo.app, err)
			continue
		}
		if !enabled {
			logrus.Errorf("%s profiler is not enabled", appInfo.app)
			continue
		}

		for filenameSuffix, ppath := range appInfo.profilerPaths {
			fs := filenameSuffix
			url := appInfo.baseURL + ppath

			wg.Add(1)

			go downloadProfilerData(url, fs, out, wg)
		}
	}

	wg.Wait()
	close(out)

	wgr.Wait()
	return files
}

func downloadProfilerData(url string, fs string, out chan pproofResult, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get(url) //nolint
	if err != nil {
		out <- pproofResult{filename: "", err: err}
		return
	}
	if resp.StatusCode != http.StatusOK {
		out <- pproofResult{filename: "", err: fmt.Errorf("cannot get profiler data(%s). Status code: %d", url, resp.StatusCode)}
		return
	}

	tmpfile, err := ioutil.TempFile("", "*_"+fs)
	if err != nil {
		out <- pproofResult{filename: "", err: errors.Wrap(err, "cannot create temp file")}
		return
	}

	if _, err := io.Copy(tmpfile, resp.Body); err != nil {
		out <- pproofResult{filename: "", err: errors.Wrap(err, "cannot write pprof file")}
		return
	}

	if err := tmpfile.Close(); err != nil {
		out <- pproofResult{filename: "", err: errors.Wrap(err, "cannot close pprof file")}
		return
	}

	out <- pproofResult{filename: tmpfile.Name(), err: nil}
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
	SummaryC.Flag("skip-server", "Skip fetching logs.zip from PMM Server").BoolVar(&Summary.SkipServer)
	SummaryC.Flag("pprof", "Include proflier information in the logs").BoolVar(&Summary.Pproof)
}
