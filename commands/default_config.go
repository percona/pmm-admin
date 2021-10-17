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
