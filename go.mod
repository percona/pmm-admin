module github.com/percona/pmm-admin

go 1.16

replace gopkg.in/alecthomas/kingpin.v2 => github.com/Percona-Lab/kingpin v2.2.6-percona+incompatible

require (
	github.com/AlekSi/pointer v1.1.0
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751 // indirect
	github.com/alecthomas/units v0.0.0-20190924025748-f65c72e2690d
	github.com/go-openapi/runtime v0.19.20
	github.com/go-openapi/spec v0.19.8 // indirect
	github.com/go-openapi/strfmt v0.19.5 // indirect
	github.com/percona/pmm v0.0.0-20210723112106-a68b7874b109
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.6.0
	github.com/stretchr/testify v1.6.1
	golang.org/x/sys v0.0.0-20200722175500-76b94024e4b6
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
)
