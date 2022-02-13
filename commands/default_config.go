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
	"io/ioutil"
	"os"
)

func DefaultConfig(val string) (f *os.File, cleanup func(), e error) {
	file, err := ioutil.TempFile("", "test-pmm-admin-defaults-*.cnf")
	if err != nil {
		return nil, nil, err
	}
	f = file
	cleanup = func() {
		_ = os.Remove(file.Name())
	}

	if _, err := file.WriteString(val); err != nil {
		return nil, nil, err
	}
	if err := file.Sync(); err != nil {
		return nil, nil, err
	}

	return
}
