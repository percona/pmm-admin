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
	"context"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/percona/pmm/api/agentlocalpb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/percona/pmm-admin/agentlocal"
)

func TestSummary(t *testing.T) {
	agentlocal.SetTransport(context.TODO(), true, agentlocal.DefaultPMMAgentListenPort)

	f, err := ioutil.TempFile("", "pmm-admin-test-summary")
	require.NoError(t, err)
	filename := f.Name()
	t.Log(filename)

	defer os.Remove(filename) //nolint:errcheck
	assert.NoError(t, f.Close())

	t.Run("Summary default", func(t *testing.T) {
		cmd := &summaryCommand{
			Filename: filename,
		}
		res, err := cmd.Run()
		require.NoError(t, err)
		expected := &summaryResult{
			Filename: filename,
		}
		assert.Equal(t, expected, res)
	})

	t.Run("Summary skip server", func(t *testing.T) {
		cmd := &summaryCommand{
			Filename:   filename,
			SkipServer: true,
		}
		res, err := cmd.Run()
		require.NoError(t, err)
		expected := &summaryResult{
			Filename: filename,
		}
		assert.Equal(t, expected, res)
	})

	t.Run("Summary pprof", func(t *testing.T) {
		if os.Getenv("DEVCONTAINER") == "" {
			t.Skip("can be tested only inside devcontainer")
		}

		cmd := &summaryCommand{
			Filename:   filename,
			SkipServer: true,
			Pprof:      true,
		}
		res, err := cmd.Run()
		require.NoError(t, err)
		expected := &summaryResult{
			Filename: filename,
		}

		// Check there is a pprof dir with data inside the zip file
		reader, err := zip.OpenReader(filename)
		assert.NoError(t, err)
		assert.Equal(t, expected, res)

		hasPprofDir := false

		for _, file := range reader.File {
			if filepath.Dir(file.Name) == "pprof" {
				hasPprofDir = true
				break
			}
		}

		assert.True(t, hasPprofDir)
	})

	t.Run("Summary - test process_exec_path", func(t *testing.T) {
		t.Parallel()
		if os.Getenv("DEVCONTAINER") == "" {
			t.Skip("can be tested only inside devcontainer")
		}

		cmd := &summaryCommand{
			Filename:   filename,
			Pprof:      false,
			SkipServer: true,
		}
		res, err := cmd.Run()
		require.NoError(t, err)
		expected := &summaryResult{
			Filename: filename,
		}

		// Check there is a pprof dir with data inside the zip file
		reader, err := zip.OpenReader(filename)
		assert.NoError(t, err)
		assert.Equal(t, expected, res)

		hasStatusFile := false
		var statusFile *zip.File

		for _, file := range reader.File {
			if file.Name == "client/status.json" {
				statusFile = file
				hasStatusFile = true

				break
			}
		}

		assert.True(t, hasStatusFile)

		jsonFile, err := statusFile.Open()
		assert.NoError(t, err)
		defer assert.NoError(t, jsonFile.Close())
		jsonByteValueFile, _ := ioutil.ReadAll(jsonFile)

		var status agentlocalpb.StatusResponse
		err = json.Unmarshal(jsonByteValueFile, &status)
		assert.NoError(t, err)

		agentHasProcessExecPath := false
		for _, agentInfo := range status.AgentsInfo {
			if agentInfo.ProcessExecPath != "" {
				agentHasProcessExecPath = true

				break
			}
		}

		assert.True(t, agentHasProcessExecPath)
	})
}
