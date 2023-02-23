# PREAMBLE
#//////////////////////////////////////////////////////////////////////////////
#
SHELL := bash

# VARIABLES, CONFIG, & SETTINGS
#//////////////////////////////////////////////////////////////////////////////
#
BINDIR       := $(CURDIR)/bin
BINNAME      ?= go-oscal
INSTALL_PATH ?= /usr/local/bin
OSCAL_COMPONENT_SCHEMA_FILE := testdata/schema/oscal_component_schema.json
OSCAL_SSP_SCHEMA_FILE := testdata/schema/oscal_ssp_schema.json
OSCAL_ASSESSMENT_PLAN_SCHEMA_FILE := testdata/schema/oscal_assessment-plan_schema.json

# Git vars
GIT_COMMIT = $(shell git rev-parse HEAD)
GIT_SHA    = $(shell git rev-parse --short HEAD)
GIT_TAG    = $(shell git describe --tags --abbrev=0 --exact-match 2>/dev/null)

# Go CLI options
PKG         := ./...
TAGS        :=
TESTS       := .
TESTFLAGS   := -race -v -failfast
LDFLAGS     := -w -s
GOFLAGS     :=
CGO_ENABLED ?= 0

# Allows us to set VERSION from the command line.
# Otherwise, if BINARY_VERSION is not set, use the current git tag.
ifdef VERSION
	BINARY_VERSION = $(VERSION)
endif
BINARY_VERSION ?= ${GIT_TAG}

# Project sources.
SRC := $(shell find . -type f -name '*.go' -print) go.mod go.sum

# TASKS
#//////////////////////////////////////////////////////////////////////////////
#
.PHONY: all
all: clean build test generate-file

.PHONY: help
help: ## Show this help message.
	@grep -E '^[a-zA-Z_/-]+:.*?## .*$$' $(MAKEFILE_LIST) \
		| sort \
		| awk 'BEGIN {FS = ":.*?## "; printf "\nUsage:\n"}; {printf "  %-15s %s\n", $$1, $$2}'
	@echo

.PHONY: clean
clean: ## Remove generated artifacts.
	go clean
	rm -rf $(BINDIR)

.PHONY: build
build: $(BINDIR)/$(BINNAME) ## Build the project.

$(BINDIR)/$(BINNAME): $(SRC)
	CGO_ENABLED=$(CGO_ENABLED) go build $(GOFLAGS) -trimpath -tags '$(TAGS)' -ldflags '$(LDFLAGS)' -o '$(BINDIR)/$(BINNAME)' .

.PHONY: generate-file
generate-file: clean build ## Generate Go structs from OSCAL JSON schema and output to 'internal/oscal/types.go'
	$(BINDIR)/$(BINNAME) -f $(OSCAL_COMPONENT_SCHEMA_FILE) -f $(OSCAL_SSP_SCHEMA_FILE) -o types.go --tags json,yaml

.PHONY: generate-stdout
generate-stdout: clean build ## Generate Go structs from OSCAL JSON schema and output to stdout
	$(BINDIR)/$(BINNAME) -f $(OSCAL_COMPONENT_SCHEMA_FILE) -f $(OSCAL_SSP_SCHEMA_FILE) --tags json,yaml

.PHONY: test
test: build ## Run automated tests.
	go test $(GOFLAGS) -run $(TESTS) $(PKG) $(TESTFLAGS)

.PHONY: run-main
run-main: ## useful for running the main.go file without having to compile
	go run main.go -f $(OSCAL_COMPONENT_SCHEMA_FILE) -f $(OSCAL_SSP_SCHEMA_FILE) --tags json,yaml

.PHONY: install
install: ## Install binary to $INSTALL_PATH.
	@install "$(BINDIR)/$(BINNAME)" "$(INSTALL_PATH)/$(BINNAME)"
