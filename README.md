# GoPlate

A comprehensive Go-based REST API boilerplate with Fiber, GORM, powerful console commands, background jobs, and task scheduling.

## Quick Start

```bash
# Clone and setup
git clone https://github.com/sheenazien8/goplate.git
cd goplate
go mod tidy

# Configure
cp .env.example .env
# Edit .env with your database settings

# Run migrations and start (using console commands)
go run main.go console db:up
make dev

# Or use traditional make commands
# make db-up
# make dev
```

## Key Commands

### Console Commands (Recommended)
```bash
go run main.go console list              # List all available commands
go run main.go console db:up             # Run database migrations
go run main.go console make:model User   # Generate new model
go run main.go console make:dto UserDto  # Generate new DTO
```

### Make Commands
```bash
make dev          # Development server with hot reload
make build        # Build application
make test         # Run tests
```

## Documentation

Complete documentation is available in the [docs/](docs/) directory:

- [Installation](docs/installation.md)
- [Quick Start](docs/quick-start.md)
- [Console Commands](docs/console-commands.md) - **New!** Powerful development tools
- [Configuration](docs/configuration.md)
- [Database](docs/database.md)
- [API Reference](docs/api-reference.md)
- [Background Tasks](docs/background-tasks.md)

## Requirements

- Go 1.22.1+
- MySQL 8.0+ or PostgreSQL 13+

## License

MIT License
