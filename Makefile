include .env

ifndef $(GOPATH)
    GOPATH=$(shell go env GOPATH)
    export GOPATH
endif

MIGRATION_FOLDER=$(CURDIR)/internal/app/migrations

ifeq ($(POSTGRES_URI),)
	POSTGRES_URI := user=$(POSTGRES_USER) password=$(POSTGRES_PASSWORD) dbname=$(POSTGRES_DB) host=$(POSTGRES_HOST) port=$(POSTGRES_PORT) sslmode=disable
endif

run:
	go run cmd/hw3/main.go

migration-create:
	$(GOPATH)/bin/goose -dir "$(MIGRATION_FOLDER)" create "$(name)" sql

migration-up:
	$(GOPATH)/bin/goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_URI)" up

migration-down:
	$(GOPATH)/bin/goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_URI)" down

migration-status:
	$(GOPATH)/bin/goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_URI)" status
