PROJECTNAME := $(shell basename "$(PWD)")
PKG_LIST := $(shell go list ./... | grep -v /vendor/)

## build: Build the server binary to execute
build:
	go build -o ${PROJECTNAME} ./cmd/server

## run: Run the server
run:
	go run ./cmd/server

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo