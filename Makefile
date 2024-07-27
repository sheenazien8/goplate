build:
	go build -o server main.go

run: build
	./server

watch:
	reflex -s -r '\.go$$' make run

MIGRATION_DIR=./db/migrations
DBMATE=dbmate

ifeq (,$(wildcard .env))
    $(error .env file not found)
endif
include .env
export $(shell sed 's/=.*//' .env)

export DB_URL := $(DB_DRIVER)://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable

.PHONY: all
all: migrate

.PHONY: install
install:
	@echo "Installing dbmate..."
	npm install -g dotenv-cli
	go install github.com/amacneil/dbmate@latest

.PHONY: create-migration
create-migration:
	@read -p "Enter migration name: " name; \
    dotenv -e .env -- $(DBMATE) -d $(MIGRATION_DIR) new $$name

.PHONY: migrate
migrate:
	$(DBMATE) -d $(MIGRATION_DIR) --url $(DB_URL) up

.PHONY: rollback
rollback:
	$(DBMATE) -d $(MIGRATION_DIR) --url $(DB_URL) down

.PHONY: create
create:
	$(DBMATE) -d $(MIGRATION_DIR) --url $(DB_URL) create

.PHONY: drop
drop:
	$(DBMATE) -d $(MIGRATION_DIR) --url $(DB_URL) drop

.PHONY: status
status:
	$(DBMATE) -d $(MIGRATION_DIR) --url $(DB_URL) status

.PHONY: dump
dump:
	$(DBMATE) -d $(MIGRATION_DIR) --url $(DB_URL) dump

.PHONY: load
load:
	$(DBMATE) -d $(MIGRATION_DIR) --url $(DB_URL) load

.PHONY: model
model:
	@read -p "Enter model name: " model_name; \
	./scripts/generate_model.sh $$model_name
