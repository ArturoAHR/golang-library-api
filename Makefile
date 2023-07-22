SHELL := /bin/bash
include .env
export

install: env
# Setup Migrations
	go install github.com/pressly/goose/v3/cmd/goose@latest
	@goose -dir $(MIGRATIONS_DIR) postgres "host=$(DATABASE_HOST) user=$(DATABASE_USER) password=$(DATABASE_PASSWORD) dbname=$(library-api) sslmode=disable" status
	goose -dir $(MIGRATIONS_DIR) up

# Setup Live Reloading
	go install github.com/cosmtrek/air@latest
	air -c .air.toml

create-migration: env
	goose -dir $(MIGRATIONS_DIR) create $(name)

env:
	@source .env