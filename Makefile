VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')
GOBIN ?= $(GOPATH)/bin
STATIK = $(GOBIN)/statik
SHASUM := $(shell which sha256sum)
PKG_LIST := $(shell go list ./...)
verbosity = 2

build_tags =

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=certik \
		  -X github.com/cosmos/cosmos-sdk/version.ServerName=certikd \
		  -X github.com/cosmos/cosmos-sdk/version.ClientName=certikcli \
		  -X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
		  -X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \

build_tags := $(strip $(build_tags))
ldflags := $(strip $(ldflags))

BUILD_FLAGS := -tags "$(build_tags)" -ldflags '$(ldflags)'

export GO111MODULE = on

all: install release lint test

install: go.sum
	GO111MODULE=on go install $(BUILD_FLAGS) ./cmd/certikd
	GO111MODULE=on go install $(BUILD_FLAGS) ./cmd/certikcli

update-swagger-docs: statik
	$(GOBIN)/statik -src=client/lcd/swagger-ui -dest=client/lcd -f -m
	@if [ -n "$(git status --porcelain)" ]; then \
        echo "\033[91mSwagger docs are out of sync!!!\033[0m";\
        exit 1;\
    else \
    	echo "\033[92mSwagger docs are in sync\033[0m";\
    fi
.PHONY: update-swagger-docs statik-clean

statik: $(STATIK)
$(STATIK):
	@echo "Installing statik..."
	@(cd /tmp && go get github.com/rakyll/statik@v0.1.6)

statik-clean:
	rm -f $(STATIK)

go.sum: go.mod
	@echo "--> Ensure dependencies have not been modified"
	GO111MODULE=on go mod verify

release: go.sum
	GOOS=linux go build $(BUILD_FLAGS) -o build/certikcli ./cmd/certikcli
	GOOS=linux go build $(BUILD_FLAGS) -o build/certikd ./cmd/certikd
	GOOS=windows go build $(BUILD_FLAGS) -o build/certikcli.exe ./cmd/certikcli
	GOOS=windows go build $(BUILD_FLAGS) -o build/certikd.exe ./cmd/certikd
	GOOS=darwin go build $(BUILD_FLAGS) -o build/certikcli-macos ./cmd/certikcli
	GOOS=darwin go build $(BUILD_FLAGS) -o build/certikd-macos ./cmd/certikd

release32: go.sum
	GOOS=linux GOARCH=386 go build $(BUILD_FLAGS) -o certikcli ./cmd/certikcli
	GOOS=linux GOARCH=386 go build $(BUILD_FLAGS) -o certikd ./cmd/certikd

tidy:
	@gofmt -s -w .
	@go mod tidy

lint: tidy
	@GO111MODULE=on golangci-lint run --config .github/.golangci.yml

test: tidy
	@GO111MODULE=on go test ${PKG_LIST}

coverage.out: tidy
	@GO111MODULE=on go test -short -coverprofile=coverage.out -covermode=atomic ${PKG_LIST}

cov: coverage.out
	@GO111MODULE=on go tool cover -func $<

coverage: coverage.out
	@GO111MODULE=on go tool cover -html $<

image: Dockerfile Dockerfile.update
	@docker rmi -f shentu-base -f shentu
	@docker build -t shentu-base -t shentu . -f Dockerfile

image.update: Dockerfile.update
	@docker rmi -f shentu
	@docker build -t shentu . -f Dockerfile.update

include .env

localnet: localnet.down image.update docker-compose.yml ./launch/localnet_client_setup.sh
	@$(RM) -r ${LOCALNET_ROOT}
	@docker run --volume $(abspath ${LOCALNET_ROOT}):/root --workdir /root -it shentu certikd testnet --v 4 --output-dir /root --server-ip-address ${LOCALNET_START_IP} --chain-id certikchain
	@docker-compose up -d
	@docker exec $(shell basename $(CURDIR))_client_1 bash /shentu/launch/localnet_client_setup.sh

localnet.client:
	@docker exec -it $(shell basename $(CURDIR))_client_1 bash

localnet.both: localnet localnet.client

localnet.down:
	@docker-compose down --remove-orphans

.PHONY: all install release release32 fix lint test cov coverage coverage.out image image.update localnet localnet.client localnet.both localnet.down

sim-full-app:
	@echo "1/4 full app simulation..."
	@go test ./app -run TestFullAppSimulation -Enabled -Commit -NumBlocks=10 -BlockSize=200 -Seed 4 -v -timeout 5m

sim-determinism:
	@echo "\n2/4 determinism simulation..."
	@go test ./app -run TestAppStateDeterminism -Enabled -Commit -NumBlocks=10 -BlockSize=200 -Seed 4 -v -timeout 5m

sim-after-import:
	@echo "\n3/4 import simulation..."
	@go test ./app -run TestAppSimulationAfterImport -Enabled -Commit -NumBlocks=10 -BlockSize=200 -Seed 4 -v -timeout 5m

sim-import-export:
	@echo "\n4/4 import & export simulation..."
	@go test ./app -run TestAppImportExport -Enabled -Commit -NumBlocks=10 -BlockSize=200 -Seed 4 -v -timeout 5m

sim: sim-full-app sim-determinism sim-after-import sim-import-export
