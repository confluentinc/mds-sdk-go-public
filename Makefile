ALL_SRC               := $(shell find . -name "*.go" | grep -v -e vendor)
GOLANGCI_LINT_VERSION := 1.44.0

GIT_REMOTE_NAME ?= origin
MASTER_BRANCH   ?= master
RELEASE_BRANCH  ?= master
DEFAULT_BUMP := patch
include ./semver.mk
include ./mk-include/cc-begin.mk
include ./mk-include/cc-vault.mk
include ./mk-include/cc-end.mk

export GO111MODULE=on
export GONOSUMDB=github.com/confluentinc,github.com/golangci/go-misc

.PHONY: deps
deps:
	cd mdsv1; go mod download; go install github.com/golangci/golangci-lint/cmd/golangci-lint@v$(GOLANGCI_LINT_VERSION)

.PHONY: lint
lint:
	cd mdsv1; golangci-lint run --disable unused,deadcode,staticcheck,errcheck --timeout=5m

.PHONY: test
test:
	cd mdsv1; go test -race -cover

.PHONY: sdk-set-bumped-version
sdk-set-bumped-version:
	echo "package mds\n\nconst SDKVersion = \"$(BUMPED_VERSION)\"" > version.go
	git add version.go

.PHONY: release
release: sdk-set-bumped-version get-release-image commit-release tag-release

.PHONY: release-ci
release-ci:
ifeq ($(CI),true)
ifneq ($(RELEASE_BRANCH),$(_empty))
	make release
endif
else
	true
endif
