BASE_PATH = $(shell pwd)
BIN_PATH := $(BASE_PATH)/bin

export PATH := $(BIN_PATH):$(PATH)

help:                           ## Display this help message.
	@echo "Please use \`make <target>\` where <target> is one of:"
	@grep '^[a-zA-Z]' $(MAKEFILE_LIST) | \
		awk -F ':.*?## ' 'NF==2 {printf "  %-26s%s\n", $$1, $$2}'

# `cut` is used to remove first `v` from `git describe` output
PMM_RELEASE_PATH ?= bin
PMM_RELEASE_VERSION ?= $(shell git describe --always --dirty | cut -b2-)
PMM_RELEASE_TIMESTAMP ?= $(shell date '+%s')
PMM_RELEASE_FULLCOMMIT ?= $(shell git rev-parse HEAD)
PMM_RELEASE_BRANCH ?= $(shell git describe --always --contains --all)

LD_FLAGS = -ldflags " \
			-X 'github.com/percona/pmm/version.ProjectName=pmm-admin' \
			-X 'github.com/percona/pmm/version.Version=$(PMM_RELEASE_VERSION)' \
			-X 'github.com/percona/pmm/version.PMMVersion=$(PMM_RELEASE_VERSION)' \
			-X 'github.com/percona/pmm/version.Timestamp=$(PMM_RELEASE_TIMESTAMP)' \
			-X 'github.com/percona/pmm/version.FullCommit=$(PMM_RELEASE_FULLCOMMIT)' \
			-X 'github.com/percona/pmm/version.Branch=$(PMM_RELEASE_BRANCH)' \
			"

release:                        ## Build pmm-admin release binary.
	env CGO_ENABLED=0 go build -v $(LD_FLAGS) -o $(PMM_RELEASE_PATH)/pmm-admin

init:                           ## Installs development tools
	rm -rf ./bin
	cd tools && go generate -x -tags=tools

install:                        ## Install pmm-admin binary.
	go install $(LD_FLAGS) ./...

install-race:                   ## Install pmm-admin binary with race detector.
	go install $(LD_FLAGS) -race ./...

TEST_FLAGS ?= -timeout=20s

test: install                   ## Run tests.
	go test $(TEST_FLAGS) ./...

test-race:                      ## Run tests with race detector.
	go test $(TEST_FLAGS) -race ./...

test-cover:                     ## Run tests and collect per-package coverage information.
	go test $(TEST_FLAGS) -coverprofile=cover.out -covermode=count ./...

test-crosscover:                ## Run tests and collect cross-package coverage information.
	go test $(TEST_FLAGS) -coverprofile=crosscover.out -covermode=count -coverpkg=./... ./...

check:                          ## Run required checkers and linters.
	go run .github/check-license.go
	$(BIN_PATH)/golangci-lint run -c=.golangci.yml ./... --new-from-rev=main
	$(BIN_PATH)/go-consistent -pedantic ./...

FILES = $(shell find . -type f -name '*.go')

format:                         ## Format source code.
	$(BIN_PATH)/gofumpt -l -w $(FILES)
	$(BIN_PATH)/goimports -local github.com/percona/pmm-admin -l -w $(FILES)
	$(BIN_PATH)/gci write --Section Standard --Section Default --Section "Prefix(github.com/percona/pmm-admin)" $(FILES)

env-up:                         ## Start development environment.
	docker-compose up --force-recreate --abort-on-container-exit --renew-anon-volumes --remove-orphans

env-down:                       ## Stop development environment.
	docker-compose down --volumes --remove-orphans

ci-reviewdog:                   ## Runs reviewdog checks.
	go run .github/check-license.go
	$(BIN_PATH)/golangci-lint run -c=.golangci-required.yml --out-format=line-number | $(BIN_PATH)/reviewdog -f=golangci-lint -level=error -reporter=github-pr-check
	$(BIN_PATH)/golangci-lint run -c=.golangci.yml --out-format=line-number | $(BIN_PATH)/reviewdog -f=golangci-lint -level=error -reporter=github-pr-review
	$(BIN_PATH)/go-consistent -pedantic -exclude "tests" ./... | $(BIN_PATH)/reviewdog -f=go-consistent -name='Required go-consistent checks' -reporter=github-pr-check
