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

package inventory

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_loadCertificate(t *testing.T) {
	cert, err := ioutil.TempFile("", "cert")
	defer func() {
		err = os.Remove(cert.Name())
		assert.NoError(t, err)
	}()
	assert.NoError(t, err)
	_, err = cert.Write([]byte("cert"))
	assert.NoError(t, err)
	err = cert.Close()
	assert.NoError(t, err)

	certificate, err := loadCertificate(cert.Name())
	assert.NoError(t, err)
	assert.Equal(t, "cert", certificate)
}
