# Build all by default, even if it's not first
.DEFAULT_GOAL := all

.PHONY: all
all: tidy build

# ==============================================================================
# Build options

ROOT_PACKAGE=goer-startup

# ==============================================================================
# Includes

include scripts/make-rules/common.mk # make sure include common.mk at the first include line
include scripts/make-rules/golang.mk

# ==============================================================================
# Targets

## build: Build source code for host platform.
.PHONY: build
build:
	@$(MAKE) go.build

## tidy: Go mod tidy
.PHONY: tidy
tidy:
	@go mod tidy

## help: Show this help info.
.PHONY: help
help: Makefile
	@echo -e "\nUsage: make <TARGETS> <OPTIONS> ...\n\nTargets:"
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /'
