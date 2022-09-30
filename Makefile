PACKAGES_SIMTEST=$(shell go list ./... | grep '/simulation')
VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')
LEDGER_ENABLED ?= true
GOBIN ?= $(GOPATH)/bin
STATIK = $(GOBIN)/statik
SHASUM := $(shell which sha256sum)
PKG_LIST := $(shell go list ./...)
DOCKER := $(shell which docker)
DOCKER_BUF := $(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace bufbuild/buf
BUILDDIR ?= $(CURDIR)/build

verbosity = 2

build_tags = netgo
ifeq ($(LEDGER_ENABLED),true)
  ifeq ($(OS),Windows_NT)
    GCCEXE = $(shell where gcc.exe 2> NUL)
    ifeq ($(GCCEXE),)
      $(error gcc.exe not installed for ledger support, please install or set LEDGER_ENABLED=false)
    else
      build_tags += ledger
    endif
  else
    UNAME_S = $(shell uname -s)
    ifeq ($(UNAME_S),OpenBSD)
      $(warning OpenBSD detected, disabling ledger support (https://github.com/cosmos/cosmos-sdk/issues/1988))
    else
      GCC = $(shell command -v gcc 2> /dev/null)
      ifeq ($(GCC),)
        $(error gcc not installed for ledger support, please install or set LEDGER_ENABLED=false)
      else
        build_tags += ledger
      endif
    endif
  endif
endif

ifeq (cleveldb,$(findstring cleveldb,$(GAIA_BUILD_OPTIONS)))
  build_tags += gcc
endif
build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))

whitespace :=
whitespace += $(whitespace)
comma := ,
build_tags_comma_sep := $(subst $(whitespace),$(comma),$(build_tags))

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=shentu \
		  -X github.com/cosmos/cosmos-sdk/version.AppName=shentud \
		  -X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
		  -X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
		  -X "github.com/cosmos/cosmos-sdk/version.BuildTags=$(build_tags_comma_sep)"

build_tags := $(strip $(build_tags))
ldflags := $(strip $(ldflags))

BUILD_FLAGS := -tags "$(build_tags)" -ldflags '$(ldflags)'

# The below include contains the tools target.
include devtools/Makefile

export GO111MODULE = on

all: install release lint test-unit

install: go.sum
	go install $(BUILD_FLAGS) ./app/shentud

proto-swagger-gen:
	@./devtools/protoc-swagger-gen.sh

update-swagger-docs: statik
	$(GOBIN)/statik -src=docs/swagger -dest=docs -f -m
	@if [ -n "$(git status --porcelain)" ]; then \
    echo "\033[91mSwagger docs are out of sync!!!\033[0m";\
    exit 1;\
  else \
    echo "\033[92mSwagger docs are in sync\033[0m";\
  fi

update-cli-docs: install
	shentud --doc docs/shentud
	@perl -pi -e "s|^#* Auto generated by .*||" docs/**/*.md
	@perl -pi -e "s|$$HOME|~|" docs/**/*.md

release: go.sum
	GOOS=linux go build $(BUILD_FLAGS) -o build/shentud ./app/shentud
	GOOS=windows go build $(BUILD_FLAGS) -o build/shentud.exe ./app/shentud
	GOOS=darwin go build $(BUILD_FLAGS) -o build/shentud-macos ./app/shentud

build: go.sum
ifeq ($(OS),Windows_NT)
	go build -mod=readonly $(BUILD_FLAGS) -o build/shentud.exe ./app/shentud
else
	go build -mod=readonly $(BUILD_FLAGS) -o build/shentud ./app/shentud
endif

build-linux: go.sum
	LEDGER_ENABLED=false GOOS=linux GOARCH=amd64 $(MAKE) build

########## Tools ##########

go-mod-cache: go.sum
	@echo "--> Download go modules to local cache"
	@go mod download

go.sum: go.mod
	@echo "--> Ensure dependencies have not been modified"
	go mod verify

clean:
	#rm -rf snapcraft-local.yaml build/
	rm -rf $(BUILDDIR)/ artifacts/

distclean:
	rm -rf \
    gitian-build-darwin/ \
    gitian-build-linux/ \
    gitian-build-windows/ \
    .gitian-builder-cache/

tidy:
	@gofmt -s -w .
	@go mod tidy

lint: tidy
	@echo "--> Running linter"
	@go run github.com/golangci/golangci-lint/cmd/golangci-lint run --timeout=10m

########## Testing ##########

PACKAGES_UNIT=$(shell go list ./... | grep -v -e '/tests/e2e')
PACKAGES_E2E=$(shell go list ./... | grep '/e2e')
TEST_PACKAGES=./...
TEST_TARGETS := test-unit test-unit-cover test-race test-e2e

test-unit: ARGS=-timeout=5m -tags='norace'
test-unit: TEST_PACKAGES=$(PACKAGES_UNIT)
test-unit-cover: ARGS=-timeout=5m -tags='norace' -coverprofile=coverage.txt -covermode=atomic
test-unit-cover: TEST_PACKAGES=$(PACKAGES_UNIT)
test-race: ARGS=-timeout=5m -race
test-race: TEST_PACKAGES=$(PACKAGES_UNIT)
test-e2e: ARGS=-timeout=25m -v
test-e2e: TEST_PACKAGES=$(PACKAGES_E2E)
$(TEST_TARGETS): run-tests

run-tests:
ifneq (,$(shell which tparse 2>/dev/null))
	@echo "--> Running tests"
	@go test -mod=readonly -json $(ARGS) $(TEST_PACKAGES) | tparse
else
	@echo "--> Running tests"
	@go test -mod=readonly $(ARGS) $(TEST_PACKAGES)
endif

.PHONY: run-tests $(TEST_TARGETS)

docker-build-debug:
	@docker build -t shentuchain/shentud-e2e --build-arg IMG_TAG=debug -f e2e.Dockerfile .

# in CI.
docker-build-hermes:
	@cd tests/e2e/docker; docker build -t cosmos/hermes-e2e:latest -f hermes.Dockerfile .


image: Dockerfile Dockerfile.update
	@docker rmi -f shentu-base -f shentud
	@docker build -t shentu-base -t shentud . -f Dockerfile

image.update: Dockerfile.update
	@docker rmi -f shentud
	@docker build -t shentud . -f Dockerfile.update

include .env

###############################################################################
###                                Localnet                                 ###
###############################################################################

build-docker-shentunode:
	$(MAKE) -C networks/local

# Run a 4-node testnet locally
localnet-start: build-linux build-docker-shentunode localnet-stop
	@if ! [ -f build/node0/shentu/config/genesis.json ]; then docker run --rm -v $(CURDIR)/build:/shentu:Z shentufoundation/shentunode testnet --v 4 -o . --starting-ip-address 192.168.10.2 --keyring-backend=test ; fi
	docker-compose up -d

# Stop testnet
localnet-stop:
	docker-compose down

start-localnet-ci:
	./build/shentud init liveness --chain-id liveness --home ~/.shentud-liveness
	./build/shentud config chain-id liveness --home ~/.shentud-liveness
	./build/shentud config keyring-backend test --home ~/.shentud-liveness
	./build/shentud keys add val --home ~/.shentud-liveness
	./build/shentud add-genesis-account val 10000000000000000000000000uctk --home ~/.shentud-liveness --keyring-backend test
	./build/shentud gentx val 1000000000uctk --home ~/.shentud-liveness --chain-id liveness
	./build/shentud collect-gentxs --home ~/.shentud-liveness
	sed -i'' 's/minimum-gas-prices = ""/minimum-gas-prices = "0uatom"/' ~/.shentud-liveness/config/app.toml
	./build/shentud start --home ~/.shentud-liveness --x-crisis-skip-assert-invariants

.PHONY: start-localnet-ci

# include simulations
include sims.mk

.PHONY: all build-linux install release release32 \
	fix lint test cov coverage coverage.out image image.update \
	build-docker-shentunode localnet-start localnet-stop \
	docker-build-debug docker-build-hermes \
