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

ifeq ($(LEDGER_ENABLED),true)
  ifeq ($(OS),Windows_NT)
    GCCEXE = $(shell where gcc.exe 2> NUL)
    ifeq ($(GCCEXE),)
      $(error gcc.exe required for ledger support but not found, please install or prepend LEDGER_ENABLED=false to omit ledger support)
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
        $(error gcc required for ledger support but not found, please install or prepend LEDGER_ENABLED=false to omit ledger support)
      else
        build_tags += ledger
      endif
    endif
  endif
endif

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=certik \
		  -X github.com/cosmos/cosmos-sdk/version.ServerName=certik \
		  -X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
		  -X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
		  -X "github.com/cosmos/cosmos-sdk/version.BuildTags=$(build_tags_comma_sep)"

build_tags := $(strip $(build_tags))
ldflags := $(strip $(ldflags))

BUILD_FLAGS := -tags "$(build_tags)" -ldflags '$(ldflags)'

# The below include contains the tools target.
include devtools/Makefile

export GO111MODULE = on

all: install release lint test

install: go.sum
	go install $(BUILD_FLAGS) ./app/certik

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
	certik --doc docs/certik
	@perl -pi -e "s|^#* Auto generated by .*||" docs/**/*.md
	@perl -pi -e "s|$$HOME|~|" docs/**/*.md

release: go.sum
	GOOS=linux go build $(BUILD_FLAGS) -o build/certik ./app/certik
	GOOS=windows go build $(BUILD_FLAGS) -o build/certik.exe ./app/certik
	GOOS=darwin go build $(BUILD_FLAGS) -o build/certik-macos ./app/certik

build: go.sum
ifeq ($(OS),Windows_NT)
	go build -mod=readonly $(BUILD_FLAGS) -o build/certik.exe ./app/certik
else
	go build -mod=readonly $(BUILD_FLAGS) -o build/certik ./app/certik
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
	@GO111MODULE=on golangci-lint run --config .golangci.yml

########## Testing ##########

test: tidy
	@GO111MODULE=on go test ${PKG_LIST}

coverage.out: tidy
	@GO111MODULE=on go test -short -coverprofile=coverage.out -covermode=atomic ${PKG_LIST}

test-cov: coverage.out
	@GO111MODULE=on go tool cover -func $<

test-cov-html: coverage.out
	@GO111MODULE=on go tool cover -html $<

image: Dockerfile Dockerfile.update
	@docker rmi -f shentu-base -f shentu
	@docker build -t shentu-base -t shentu . -f Dockerfile

image.update: Dockerfile.update
	@docker rmi -f shentu
	@docker build -t shentu . -f Dockerfile.update

include .env

###############################################################################
###                                Localnet                                 ###
###############################################################################

build-docker-certiknode:
	$(MAKE) -C networks/local

# Run a 4-node testnet locally
localnet-start: build-linux build-docker-certiknode localnet-stop
	@if ! [ -f build/node0/certik/config/genesis.json ]; then docker run --rm -v $(CURDIR)/build:/certik:Z certikfoundation/certiknode testnet --v 4 -o . --starting-ip-address 192.168.10.2 --keyring-backend=test ; fi
	docker-compose up -d

# Stop testnet
localnet-stop:
	docker-compose down

# include simulations
include sims.mk

.PHONY: all build-linux install release release32 \
	fix lint test cov coverage coverage.out image image.update \
	build-docker-certiknode localnet-start localnet-stop \
