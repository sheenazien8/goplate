# GoPlate

A comprehensive Go-based REST API boilerplate with Fiber, GORM, background jobs, and task scheduling.

## Quick Start

```bash
# Clone and setup
git clone https://github.com/sheenazien8/goplate.git
cd goplate
go mod tidy

# Configure
cp .env.example .env
# Edit .env with your database settings

# Run migrations and start
make db-up
make dev
```

## Key Commands

```bash
make dev          # Development server with hot reload
make build        # Build application
make test         # Run tests
make db-create    # Create migration
make model        # Generate model
```

## Documentation

Complete documentation is available in the [docs/](docs/) directory:

- [Installation](docs/installation.md)
- [Quick Start](docs/quick-start.md)
- [Configuration](docs/configuration.md)
- [Database](docs/database.md)
- [API Reference](docs/api-reference.md)
- [Background Tasks](docs/background-tasks.md)

## Requirements

- Go 1.22.1+
- MySQL 8.0+ or PostgreSQL 13+

## License

MIT License
