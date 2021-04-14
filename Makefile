PROJECTNAME := $(shell basename "$(PWD)")
PKG_LIST := $(shell go list ./... | grep -v /vendor/)

## lint: Run the lint in projet
lint:
	golint -set_exit_status ${PKG_LIST}

## build: Build the server binary to execute
build:
	go build -o ${PROJECTNAME} ./cmd/server

## run: Run the server
run:
	go run ./cmd/server

## test: Run the test of project
test:
	go test ./... -race -coverpkg=./... -coverprofile=coverage.out

## coverage: Get coverage of all tests
coverage:
	go tool cover -func=coverage.out

## coverage-html: Generate the report in HTML of test coverage
coverage-html:
	go tool cover -html=coverage.out -o coverage.html 

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo