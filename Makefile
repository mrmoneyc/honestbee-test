GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GODEP=dep

CURR_DIR=$(shell pwd)
SRV_SRC=server
CLI_SRC=client
BUILD_OUT=bin
BUILD_SRV_BIN_NAME=tcpsrv
BUILD_CLI_BIN_NAME=tcpcli

all: clean fmt test build

test: test_server

test_server:
	@echo "--> Testing (Server)..."
	cd $(SRV_SRC); $(GOTEST) -v

build: build_server build_client

build_server:
	@echo "--> Building (Server)..."
	@mkdir -p $(BUILD_OUT)
	cd $(SRV_SRC); $(GOBUILD) -v -o ../$(BUILD_OUT)/$(BUILD_SRV_BIN_NAME)

build_client:
	@echo "--> Building (Client)..."
	@mkdir -p $(BUILD_OUT)
	cd $(CLI_SRC); $(GOBUILD) -v -o ../$(BUILD_OUT)/$(BUILD_CLI_BIN_NAME)

clean:
	@echo "--> Cleaning..."
	$(GOCLEAN)
	rm -rf $(BUILD_OUT)

fmt:
	$(GOCMD) fmt ./...
	$(GOCMD) vet ./...

.PHONY: all test test_server build build_server build_client clean fmt
