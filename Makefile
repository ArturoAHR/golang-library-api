SHELL := /bin/bash
include .env
export

install: env
# Setup Migrations
	go install github.com/pressly/goose/v3/cmd/goose@latest
	run-migrations

# Setup Live Reloading
	go install github.com/cosmtrek/air@latest
	air -c .air.toml

# Migration Commands
create-migration: env
	goose -dir $(MIGRATIONS_DIR) create $(name) sql

run-migrations: env
	goose -dir $(MIGRATIONS_DIR) postgres $(GOOSE_DSN) up

goose: env
	goose -dir $(MIGRATIONS_DIR) postgres $(GOOSE_DSN) $(command)

# Seeder Commands
run-seeder: env
	go run scripts/seeder/main.go

env:
	@source .env