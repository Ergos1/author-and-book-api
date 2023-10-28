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
