# --------------------------------------------------------------------------
# Makefile for the Fantom API GaphQL Server
#
# v0.1 (2020/03/09)  - Initial version, base API server build.
# (c) Fantom Foundation, 2020
# --------------------------------------------------------------------------

# project related vars
PROJECT := $(shell basename "$(PWD)")

# go related vars
GOBASE := $(shell pwd)
GOBIN=$(CURDIR)/build

## server: Make the API server as bin/frd
server:
	go build -o $(GOBIN)/apiserver ./cmd/apiserver

.PHONY: help
all: help
help: Makefile
	@echo
	@echo "Choose a make command in "$(PROJECT)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
