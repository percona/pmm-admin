package helpers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/percona/pmm-admin/agentlocal"
)

func GetServerVersion() (float64, error) {
	status, err := agentlocal.GetStatus(agentlocal.DoNotRequestNetworkInfo)
	if err != nil {
		return 0, err
	}

	split := strings.Split(status.ServerVersion, "-")
	split = strings.Split(split[0], ".")
	if len(split) < 3 {
		return 0, fmt.Errorf("failed to parse server version %s", status.ServerVersion)
	}

	f, err := strconv.ParseFloat(fmt.Sprintf("%s.%s%s", split[0], split[1], split[2]), 64)
	if err != nil {
		return 0, err
	}

	return f, nil
}
