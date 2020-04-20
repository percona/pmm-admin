package inventory

import (
	"testing"

	"github.com/percona/pmm/api/inventorypb/json/client/services"
	"github.com/stretchr/testify/require"
)

func TestAddServiceMySQL(t *testing.T) {
	t.Run("Address and port", func(t *testing.T) {
		res := &addServiceMySQLResult{
			Service: &services.AddMySQLServiceOKBodyMysql{
				ServiceID:      "/service_id/1",
				ServiceName:    "MySQL Service",
				NodeID:         "/node_id/1",
				Address:        "127.0.0.1",
				Port:           3306,
				Environment:    "environment",
				Cluster:        "mysql-cluster",
				ReplicationSet: "mysql-replication-set",
				CustomLabels:   map[string]string{"key": "value", "foo": "bar"},
			},
		}
		expected := `MySQL Service added.
Service ID     : /service_id/1
Service name   : MySQL Service
Node ID        : /node_id/1
Address        : 127.0.0.1
Port           : 3306
Environment    : environment
Cluster name   : mysql-cluster
Replication set: mysql-replication-set
Custom labels  : map[foo:bar key:value]
`
		require.Equal(t, expected, res.String())
	})

	t.Run("Socket", func(t *testing.T) {
		res := &addServiceMySQLResult{
			Service: &services.AddMySQLServiceOKBodyMysql{
				ServiceID:      "/service_id/1",
				ServiceName:    "MySQL Socket Service",
				NodeID:         "/node_id/1",
				Socket:         "/path/to/socket",
				Environment:    "environment",
				Cluster:        "mysql-cluster",
				ReplicationSet: "mysql-replication-set",
				CustomLabels:   map[string]string{"key": "value", "foo": "bar"},
			},
		}
		expected := `MySQL Service added.
Service ID     : /service_id/1
Service name   : MySQL Socket Service
Node ID        : /node_id/1
Socket         : /path/to/socket
Environment    : environment
Cluster name   : mysql-cluster
Replication set: mysql-replication-set
Custom labels  : map[foo:bar key:value]
`
		require.Equal(t, expected, res.String())
	})
}
