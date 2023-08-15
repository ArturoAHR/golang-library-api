SHELL := /bin/bash
include .env
export

setup: env install-tools run-migrations run-seeder
	air -c .air.toml

install-tools:
	go install github.com/cosmtrek/air@latest
	go install github.com/pressly/goose/v3/cmd/goose@latest

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