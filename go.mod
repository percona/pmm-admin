module github.com/percona/pmm-admin

go 1.15

require (
	github.com/AlekSi/pointer v1.1.0
	github.com/Percona-Lab/kingpin v2.2.7-0.20190911101335-33f03abf8b59+incompatible
	github.com/PuerkitoBio/purell v1.1.1
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/alecthomas/units v0.0.0-20190924025748-f65c72e2690d
	github.com/asaskevich/govalidator v0.0.0-20180315120708-ccb8e960c48f
	github.com/davecgh/go-spew v1.1.1
	github.com/go-openapi/analysis v0.19.6
	github.com/go-openapi/errors v0.19.6
	github.com/go-openapi/jsonpointer v0.19.3
	github.com/go-openapi/jsonreference v0.19.3
	github.com/go-openapi/loads v0.19.4
	github.com/go-openapi/runtime v0.19.19
	github.com/go-openapi/spec v0.19.8
	github.com/go-openapi/strfmt v0.19.3
	github.com/go-openapi/swag v0.19.5
	github.com/go-openapi/validate v0.19.10
	github.com/go-stack/stack v1.8.0
	github.com/go-toolsmith/astcast v1.0.0
	github.com/go-toolsmith/astequal v1.0.0
	github.com/go-toolsmith/astinfo v1.0.0
	github.com/go-toolsmith/pkgload v1.0.0
	github.com/go-toolsmith/typep v1.0.0
	github.com/kisielk/gotool v1.0.0
	github.com/konsorten/go-windows-terminal-sequences v1.0.2
	github.com/mailru/easyjson v0.7.0
	github.com/mitchellh/mapstructure v1.1.2
	github.com/percona/pmm v2.11.2-0.20201112072023-1255fe3fdb2e+incompatible
	github.com/pkg/errors v0.8.1
	github.com/pmezard/go-difflib v1.0.0
	github.com/quasilyte/go-consistent v0.0.0-20190521200055-c6f3937de18c
	github.com/sirupsen/logrus v1.6.0
	github.com/stretchr/testify v1.4.0
	go.mongodb.org/mongo-driver v1.1.3
	golang.org/x/net v0.0.0-20191112182307-2180aed22343
	golang.org/x/sys v0.0.0-20191112214154-59a1497f0cea
	golang.org/x/text v0.3.2
	golang.org/x/tools v0.0.0-20191113055240-e33b02e76616
	gopkg.in/yaml.v2 v2.2.5
)

// replace gopkg.in/alecthomas/kingpin.v2 33f03abf8b596e750f3cc9d66a2555d6cc3daa78 => github.com/Percona-Lab/kingpin v2.2.7-0.20190911101335-33f03abf8b59+incompatible
replace gopkg.in/alecthomas/kingpin.v2 => github.com/Percona-Lab/kingpin v2.2.6-percona
