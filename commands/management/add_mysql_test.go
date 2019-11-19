package management

import (
	"strings"
	"testing"

	mysql "github.com/percona/pmm/api/managementpb/json/client/my_sql"

	"github.com/stretchr/testify/assert"
)

func TestAddMySQL(t *testing.T) {
	t.Run("TablestatEnabled", func(t *testing.T) {
		res := &addMySQLResult{
			Service: &mysql.AddMySQLOKBodyService{
				ServiceID:   "/service/1",
				ServiceName: "mysql-1",
			},
			MysqldExporter: &mysql.AddMySQLOKBodyMysqldExporter{
				TablestatsGroupTableLimit: 1000,
				TablestatsGroupDisabled:   false,
			},
			TableCount: 500,
		}
		expected := strings.TrimSpace(`
MySQL Service added.
Service ID  : /service/1
Service name: mysql-1

Table statistics collection enabled.
	`)
		assert.Equal(t, expected, strings.TrimSpace(res.String()))
	})

	t.Run("TablestatDisabled", func(t *testing.T) {
		res := &addMySQLResult{
			Service: &mysql.AddMySQLOKBodyService{
				ServiceID:   "/service/1",
				ServiceName: "mysql-1",
			},
			MysqldExporter: &mysql.AddMySQLOKBodyMysqldExporter{
				TablestatsGroupTableLimit: 1000,
				TablestatsGroupDisabled:   true,
			},
			TableCount: 2000,
		}
		expected := strings.TrimSpace(`
MySQL Service added.
Service ID  : /service/1
Service name: mysql-1

Table statistics collection disabled.
	`)
		assert.Equal(t, expected, strings.TrimSpace(res.String()))
	})
}
