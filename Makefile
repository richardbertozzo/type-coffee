PROJECTNAME := $(shell basename "$(PWD)")
PKG_LIST := $(shell go list ./... | grep -v /vendor/)
DATABASE := typecoffee

## lint: Run the lint
lint:
	go vet ${PKG_LIST}

## build: Build the api binary to execute
build:
	go build -o ${PROJECTNAME} ./cmd/api

## run: Run the api
run:
	go run ./cmd/api

## test: Run the test of project
test:
	go test ./... -race -coverpkg=./... -coverprofile=coverage.out

## coverage: Get coverage of all tests
coverage:
	go tool cover -func=coverage.out

## coverage-html: Generate the report in HTML of test coverage
coverage-html:
	go tool cover -html=coverage.out -o coverage.html 

.PHONY: gen-go-openapi-code
gen-go-openapi-code:
	mkdir -p coffee/handler
	oapi-codegen --config configs/openapi/types.yml api/openapi.yaml
	oapi-codegen --config configs/openapi/server.yml api/openapi.yaml

.PHONY: gen-go-sql-code
gen-go-sql-code:
	sqlc generate --file configs/db/sqlc.yaml

.PHONY: pg-up
pg-up:
	docker run \
		--name postgres \
		-p 5432:5432 \
		-e POSTGRES_USER=root \
		-e POSTGRES_PASSWORD=secret \
		-e POSTGRES_DB=${DATABASE} \
		-d --rm postgres:14

.PHONY: pg-down
pg-down:
	docker stop postgres

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo