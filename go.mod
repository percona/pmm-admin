module github.com/percona/pmm-admin

go 1.16

replace gopkg.in/alecthomas/kingpin.v2 v2.2.7-0.20190911101335-33f03abf8b59+incompatible => github.com/Percona-Lab/kingpin v2.2.7-0.20190911101335-33f03abf8b59+incompatible

require (
	github.com/AlekSi/pointer v1.1.0
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751 // indirect
	github.com/alecthomas/units v0.0.0-20190924025748-f65c72e2690d
	github.com/go-openapi/runtime v0.19.19
	github.com/percona/pmm v2.14.1-0.20210308095546-6176fdc415f7+incompatible
	github.com/pkg/errors v0.8.1
	github.com/sirupsen/logrus v1.6.0
	github.com/stretchr/testify v1.6.1
	golang.org/x/sys v0.0.0-20200323222414-85ca7c5b95cd
	gopkg.in/alecthomas/kingpin.v2 v2.2.7-0.20190911101335-33f03abf8b59+incompatible
)
