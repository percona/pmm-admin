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
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/percona/pmm-admin/logger"
)

func init() {
	logrus.SetFormatter(new(logger.TextFormatter))
}

func TestParseRenderTemplate(t *testing.T) {
	var stderr bytes.Buffer
	logrus.SetOutput(&stderr)
	defer logrus.SetOutput(os.Stderr)

	tmpl := ParseTemplate(`{{ .Missing }}`)
	data := map[string]string{
		"foo": "bar",
	}

	assert.Panics(t, func() { RenderTemplate(tmpl, data) })

	expected := strings.TrimSpace(`
Failed to render response.
template: :1:3: executing "" at <.Missing>: map has no entry for key "Missing".
Template data: map[string]string{"foo":"bar"}.
Please report this bug.
	`) + "\n"
	assert.Equal(t, expected, stderr.String())
}

func TestParseCustomLabel(t *testing.T) {
	errWrongFormat := fmt.Errorf("wrong custom label format")
	for _, v := range []struct {
		name     string
		input    string
		expected map[string]string
		expErr   error
	}{
		{"simple label", "foo=bar", map[string]string{"foo": "bar"}, nil},
		{"two labels", "foo=bar,bar=foo", map[string]string{"foo": "bar", "bar": "foo"}, nil},
		{"no value", "foo=", nil, errWrongFormat},
		{"no key", "=foo", nil, errWrongFormat},
		{"wrong format", "foo=bar,bar+foo", nil, errWrongFormat},
		{"empty value", "", map[string]string{}, nil},
		{"PMM-4078 hyphen", "region=us-east1, mylabel=mylab-22", map[string]string{"region": "us-east1", "mylabel": "mylab-22"}, nil},
	} {
		t.Run(v.name, func(t *testing.T) {
			customLabels, err := ParseCustomLabels(v.input)
			assert.Equal(t, v.expected, customLabels)
			assert.Equal(t, v.expErr, err)
		})
	}
}

func TestReadFile(t *testing.T) {
	t.Run("Normal", func(t *testing.T) {
		cert, err := ioutil.TempFile("", "cert")
		require.NoError(t, err)
		defer func() {
			err = cert.Close()
			assert.NoError(t, err)
			err = os.Remove(cert.Name())
			assert.NoError(t, err)
		}()
		_, err = cert.Write([]byte("cert"))
		require.NoError(t, err)

		certificate, err := ReadFile(cert.Name())
		assert.NoError(t, err)
		assert.Equal(t, "cert", certificate)
	})

	t.Run("WrongPath", func(t *testing.T) {
		filepath := "not-existed-file"
		certificate, err := ReadFile(filepath)
		assert.EqualError(t, err, fmt.Sprintf("cannot load file in path %q: open not-existed-file: no such file or directory", filepath))
		assert.Empty(t, certificate)
	})

	t.Run("EmptyFilePath", func(t *testing.T) {
		certificate, err := ReadFile("")
		require.NoError(t, err)
		require.Empty(t, certificate)
	})
}
