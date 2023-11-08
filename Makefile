include .env

ifndef $(GOPATH)
    GOPATH=$(shell go env GOPATH)
    export GOPATH
endif

MIGRATION_FOLDER=$(CURDIR)/internal/app/migrations

ifeq ($(POSTGRES_URI),)
	POSTGRES_URI := user=$(POSTGRES_USER) password=$(POSTGRES_PASSWORD) dbname=$(POSTGRES_DB) host=$(POSTGRES_HOST) port=$(POSTGRES_PORT) sslmode=disable
endif

ifeq ($(POSTGRES_URI_PLAIN),)
	POSTGRES_URI_PLAIN := postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)
endif

up-deps:
	docker-compose -f deployments/psql-db/docker-compose.yml up -d 
	docker-compose -f deployments/kafka/docker-compose.yml up -d 

down-deps:
	docker-compose -f deployments/psql-db/docker-compose.yml up -d 
	docker-compose -f deployments/kafka/docker-compose.yml up -d 

run:
	go run cmd/app_http/main.go

clear-db:
	psql $(POSTGRES_URI_PLAIN) -c "TRUNCATE authors, books"

migration-create:
	$(GOPATH)/bin/goose -dir "$(MIGRATION_FOLDER)" create "$(name)" sql

migration-up:
	$(GOPATH)/bin/goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_URI)" up

migration-down:
	$(GOPATH)/bin/goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_URI)" down

migration-status:
	$(GOPATH)/bin/goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_URI)" status


test-integration:
	go test -v ./... -tags=integration

test-unit:
	go test -v ./... -tags=unit


GO_PB_PATH:="api/v1"
AUTHOR_PB_PATH:="${GO_PB_PATH}/author.proto"
BOOK_PB_PATH:="${GO_PB_PATH}/book.proto"

GO_OUT_PB_PATH:="pkg/api/grpc/v1"
AUTHOR_OUT_PB_PATH:="${GO_OUT_PB_PATH}/author"
BOOK_OUT_PB_PATH:="${GO_OUT_PB_PATH}/book"

protoc-generate:
	mkdir -p ${AUTHOR_OUT_PB_PATH}
	mkdir -p ${BOOK_OUT_PB_PATH}
	protoc -I=${GO_PB_PATH}/. --proto_path=. --go_out=${BOOK_OUT_PB_PATH} --go_opt paths=source_relative --go-grpc_out=${BOOK_OUT_PB_PATH} --go-grpc_opt paths=source_relative ${BOOK_PB_PATH} 
	protoc -I=${GO_PB_PATH}/. --proto_path=. --go_out=${AUTHOR_OUT_PB_PATH} --go_opt paths=source_relative --go-grpc_out=${AUTHOR_OUT_PB_PATH} --go-grpc_opt paths=source_relative ${AUTHOR_PB_PATH} 