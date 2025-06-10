# GoPlate Makefile - Development and Database Management Tools

# Default target
.DEFAULT_GOAL := help

# Build the application
build:
	@echo "🔨 Building application..."
	go build -o server main.go

# Run the application (builds first)
run: build
	@echo "🚀 Starting server..."
	./server

# Watch for changes and auto-reload (requires reflex: go install github.com/cespare/reflex@latest)
watch:
	@echo "👀 Watching for changes..."
	reflex -s -r '\.go$$' make run

# Development server with hot reload
dev: watch

# Clean build artifacts
clean:
	@echo "🧹 Cleaning build artifacts..."
	rm -f server

# Format Go code
fmt:
	@echo "✨ Formatting code..."
	go fmt ./...

# Run tests
test:
	@echo "🧪 Running tests..."
	go test ./...

# Run tests with coverage
test-coverage:
	@echo "📊 Running tests with coverage..."
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Install development dependencies
install-deps:
	@echo "📦 Installing development dependencies..."
	go install github.com/cespare/reflex@latest
	go install github.com/amacneil/dbmate@latest
	npm install -g dotenv-cli

# Tidy go modules
tidy:
	@echo "🔧 Tidying go modules..."
	go mod tidy

# Show help
help:
	@echo "🚀 GoPlate Development Commands"
	@echo ""
	@echo "📋 Available commands:"
	@echo "  build                       Build the application"
	@echo "  run                         Build and run the application"
	@echo "  dev/watch                   Start development server with hot reload"
	@echo "  clean                       Clean build artifacts"
	@echo "  fmt                         Format Go code"
	@echo "  test                        Run tests"
	@echo "  test-coverage               Run tests with coverage report"
	@echo "  install-deps                Install development dependencies"
	@echo "  tidy                        Tidy go modules"
	@echo ""
	@echo "🗄️  Database commands:"
	@echo "  db-create                   Create a new migration file (interactive)"
	@echo "  db-up                       Run pending migrations"
	@echo "  db-down                     Rollback last migration"
	@echo "  db-status                   Show migration status"
	@echo "  db-reset                    Drop and recreate database"
	@echo "  db-fresh                    Fresh migration (reset + migrate)"
	@echo "  db-dump                     Dump database schema"
	@echo "  db-load                     Load database schema"
	@echo "  db-connect                  Connect to database interactive shell"
	@echo "  db-seeder-create            Create a new seeder file"
	@echo "  db-seeder-run               Run all seeders"
	@echo "  db-seeder-run filename      Run seeder by filename"
	@echo "  db-help                     Show help for database commands"
	@echo ""
	@echo "🏗️  Code generation:"
	@echo "  model                       Generate a new model"
	@echo "  dto                         Generate a new DTO"
	@echo "  cron                        Generate a new cron scheduler file"
	@echo "  job                         Generate a new listener queue job file"
	@echo ""
	@echo "💡 Examples:"
	@echo "  make dev                    # Start development server"
	@echo "  make db-create              # Create new migration"
	@echo "  make model                  # Generate new model"
	@echo "  make test-coverage          # Run tests with coverage"

# Database migration commands (using ./scripts/migrate.sh)
.PHONY: db-create
db-create:
	@read -p "Enter migration name: " name; \
	./scripts/migrate.sh create $$name

.PHONY: db-up
db-up:
	@./scripts/migrate.sh up

.PHONY: db-down
db-down:
	@./scripts/migrate.sh down

.PHONY: db-status
db-status:
	@./scripts/migrate.sh status

.PHONY: db-reset
db-reset:
	@./scripts/migrate.sh reset

.PHONY: db-fresh
db-fresh:
	@./scripts/migrate.sh fresh

.PHONY: db-dump
db-dump:
	@./scripts/migrate.sh dump

.PHONY: db-load
db-load:
	@./scripts/migrate.sh load

.PHONY: db-version
db-version:
	@./scripts/migrate.sh version

.PHONY: db-connect
db-connect:
	@./scripts/migrate.sh connect

.PHONY: db-seeder-create
db-seeder-create:
	@read -p "Enter seeder name: " name; \
	./scripts/migrate.sh seeder-create $$name

.PHONY: db-seeder-run
db-seeder-run:
	@./scripts/migrate.sh seeder-run $(file)

.PHONY: db-help
db-help:
	@./scripts/migrate.sh help

# Legacy aliases (deprecated - use db-* commands)
.PHONY: create-migration migrate rollback status
create-migration: db-create
migrate: db-up
rollback: db-down
status: db-status

.PHONY: model
model:
	@read -p "Enter model name: " model_name; \
	./scripts/generate_model.sh $$model_name

.PHONY: dto
dto:
	@read -p "Enter dto name: " dto_name; \
	./scripts/generate_dto.sh $$dto_name

.PHONY: cron
cron:
	@read -p "Enter cron file name: " cron_file; \
	./scripts/generate_cronfile.sh $$cron_file

.PHONY: job
job:
	@read -p "Enter job file name: " job_file; \
	./scripts/generate_job.sh $$job_file
