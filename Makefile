NAME := systracer
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := -X 'main.revision=$(REVISION)'
DIR_BIN := ./bin

.DEFAULT_GOAL := help

.PHONY: setup
setup:  ## Setup for required tools.
	go get golang.org/x/tools/cmd/goimports

.PHONY: fmt
fmt: ## Formatting source codes.
	@goimports -w .

.PHONY: lint
lint: ## Run go vet.
	@go vet ./...

.PHONY: build
build: _example/linux-386/main.go  ## Build a binary.
	mkdir -p $(DIR_BIN)
	gcc ./_example/hello.c -o $(DIR_BIN)/hello -Wall -O0
	GOOS=linux GOARCH=386 go build -o $(DIR_BIN)/systracer-linux-386 -ldflags "$(LDFLAGS)" ./_example/linux-386/main.go

.PHONY: help
help: ## Show help text
	@echo "Commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "    \033[36m%-20s\033[0m %s\n", $$1, $$2}'