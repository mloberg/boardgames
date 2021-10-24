PKG_LIST := $(shell go list ./... | grep -v /vendor/)
USERNAME ?= mloberg

help:
	@echo "+ $@"
	@grep -hE '(^[a-zA-Z0-9\._-]+:.*?##.*$$)|(^##)' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}{printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}' | sed -e 's/\[32m## /[33m/'
.PHONY: help

##
## Site
## ---------------------------------------------------------------------------

serve: ## Serve site
	@echo "+ $@"
	@hugo server
.PHONY: serve

##
## Index
## ---------------------------------------------------------------------------

build: bin/bg ## Build the executable
.PHONY: build

bin/bg: **/*.go
	@go build -o bin/bg

setup: bin/bg ## Setup index
	@echo "+ $@"
	@bin/bg setup
.PHONY: setup

index: bin/bg ## Load from BoardGameGeek into MeiliSearch
	@echo "+ $@"
	@bin/bg index $(USERNAME)
.PHONY: index

##
## Development
## ---------------------------------------------------------------------------

mod: ## Make sure go.mod is up to date
	@echo "+ $@"
	@go mod tidy
.PHONY: mod

lint: ## Lint Go code
	@echo "+ $@"
	@golangci-lint run
.PHONY: lint

generate: ## Autogenerate docs and resources
	@echo "+ $@"
	@go generate ${PKG_LIST}
.PHONY: generate
