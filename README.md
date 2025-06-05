# 🚀 GoPlate

A modern, production-ready Go boilerplate for building REST APIs with best practices built-in.

## ✨ Features

- **🔥 Fiber Framework** - Fast HTTP framework built on top of Fasthttp
- **📊 Structured Logging** - JSON logging with Logrus and file rotation
- **🗄️ Database Integration** - GORM with MySQL/PostgreSQL support
- **🔐 Authentication Ready** - JWT middleware and user management
- **📝 Request Validation** - Built-in validation with go-playground/validator
- **🛠️ Developer Tools** - Hot reload, migrations, and seeders
- **📦 Clean Architecture** - Organized project structure following Go conventions
- **🔧 Environment Config** - Environment-based configuration management

## 🚀 Quick Start

### Installation

Install the GoPlate CLI tool:

```bash
curl -sSL https://raw.githubusercontent.com/sheenazien8/goplate/master/install.sh | bash
```

Or install manually:
```bash
go install github.com/sheenazien8/goplate/cmd/goplate@latest
```

### Create a New Project

```bash
goplate my-awesome-api
cd my-awesome-api
go mod tidy
```

### Environment Setup

```bash
cp .env.example .env
```

Edit `.env` with your configuration:
```env
APP_NAME=MyAwesomeAPI
APP_ENV=local
APP_DEBUG=true
APP_URL=http://localhost
APP_PORT=8080
APP_SCREET=your-secret-key-here

DB_CONNECTION=mysql
DB_HOST=localhost
DB_PORT=3306
DB_DATABASE=my_database
DB_USERNAME=root
DB_PASSWORD=your_password
```

### Run the Application

```bash
go run main.go
```

Your API will be available at `http://localhost:8080`

## 📁 Project Structure

```
├── cmd/                    # CLI tools and installers
├── config/                 # Configuration management
├── db/                     # Database connection and migrations
│   ├── migrations/         # SQL migration files
│   └── seeders/           # Database seeders
├── env/                    # Environment variable utilities
├── logs/                   # Logging utilities
├── pkg/                    # Main application packages
│   ├── controllers/        # HTTP handlers
│   ├── models/            # Database models
│   └── utils/             # Utility functions
├── router/                 # Route definitions
├── scripts/               # Helper scripts
├── .env.example           # Environment variables template
├── go.mod                 # Go module file
├── main.go                # Application entry point
└── Makefile              # Build and development tasks
```

## 🛠️ Development

### Available Commands

GoPlate includes a comprehensive Makefile with development tools. Run `make help` to see all available commands:

#### 🚀 Development Commands
```bash
make help                   # Show all available commands with descriptions
make dev                    # Start development server with hot reload
make run                    # Build and run the application
make build                  # Build the application binary
make clean                  # Clean build artifacts
make fmt                    # Format Go code
make test                   # Run tests
make test-coverage          # Run tests with coverage report (generates coverage.html)
make tidy                   # Tidy go modules
make install-deps           # Install development dependencies
```

#### 🗄️ Database Commands
```bash
make install-deps          # Install database tools (dbmate, dotenv-cli)
make db-create             # Create a new migration file (interactive)
make db-up                 # Run pending migrations
make db-down               # Rollback last migration
make db-status             # Show migration status
make db-reset              # Drop and recreate database
make db-fresh              # Fresh migration (reset + migrate)
make db-dump               # Dump database schema
make db-load               # Load database schema
make db-version            # Show current migration version
make db-connect            # Connect to database interactive shell
make db-seeder-create      # Create a new seeder file
make db-seeder-run         # Run all seeders
```

#### 🏗️ Code Generation Commands
```bash
make model                 # Generate a new model (interactive)
make dto                   # Generate a new DTO (interactive)
```

#### 📝 Usage Examples
```bash
# Start development with hot reload
make dev

# Create a new migration
make db-create
# Enter migration name: create_products_table

# Generate a new model
make model
# Enter model name: Product

# Run tests with coverage
make test-coverage
# Opens coverage.html in browser

# Database operations
make db-up                 # Apply migrations
make db-status             # Check migration status
make db-fresh              # Reset database and run all migrations
make db-connect            # Connect to database shell
```

### Generate DTOs

Use the provided script to generate Data Transfer Objects:

```bash
./scripts/generate_dto.sh UserCreateDTO
```

### Generate Models

Create new database models:

```bash
./scripts/generate_model.sh Product
```

### Database Migrations

GoPlate includes a powerful migration system using the `./scripts/migrate.sh` script:

```bash
# Create a new migration
./scripts/migrate.sh create create_products_table

# Run pending migrations
./scripts/migrate.sh up

# Check migration status
./scripts/migrate.sh status

# Rollback last migration
./scripts/migrate.sh down

# Reset database (drop + recreate)
./scripts/migrate.sh reset

# Fresh migration (reset + run all migrations)
./scripts/migrate.sh fresh

# Show current migration version
./scripts/migrate.sh version

# Connect to database interactive shell
./scripts/migrate.sh connect
```

The migration script supports both MySQL and PostgreSQL databases and includes:
- ✅ Automatic database URL building from `.env` variables
- 🛡️ Safety confirmations for destructive operations
- 🎨 Colored output for better readability
- 📋 Comprehensive status reporting
- 🔌 Interactive database shell connection (like `php artisan db`)

## 🗄️ Database Support

GoPlate supports multiple database drivers:

- **MySQL** (default)
- **PostgreSQL**

Configure your database in the `.env` file:

```env
# For MySQL
DB_CONNECTION=mysql
DB_HOST=localhost
DB_PORT=3306

# For PostgreSQL
DB_CONNECTION=postgres
DB_HOST=localhost
DB_PORT=5432
```

## 📋 API Examples

### User Registration

```bash
curl -X POST http://localhost:8080/api/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "password123"
  }'
```

### User Login

```bash
curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "password123"
  }'
```

## 📊 Logging

GoPlate uses structured JSON logging with automatic file rotation:

```go
import "github.com/yourname/yourproject/logs"

// Simple logging
logs.Info("Server started")
logs.Error("Database connection failed")

// Structured logging
logs.WithFields(logrus.Fields{
    "user_id": 123,
    "action": "login",
}).Info("User logged in")

// Formatted logging
logs.Infof("Server started on port %d", port)
```

Logs are automatically rotated daily and kept for 7 days.

## 🔧 Configuration

GoPlate supports environment-based configuration. All settings can be overridden via environment variables:

- `APP_NAME` - Application name
- `APP_ENV` - Environment (local, staging, production)
- `APP_DEBUG` - Debug mode (true/false)
- `APP_PORT` - Server port
- `APP_SCREET` - JWT secret key
- `DB_CONNECTION` - Database driver (mysql/postgres)
- `DB_HOST` - Database host
- `DB_PORT` - Database port
- `DB_DATABASE` - Database name
- `DB_USERNAME` - Database username
- `DB_PASSWORD` - Database password

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- [Fiber](https://github.com/gofiber/fiber) - Express inspired web framework
- [GORM](https://gorm.io/) - The fantastic ORM library for Golang
- [Logrus](https://github.com/sirupsen/logrus) - Structured logger for Go
- [Validator](https://github.com/go-playground/validator) - Go Struct and Field validation
- [Db Mate](https://github.com/amacneil/dbmate) - A lightweight, framework-agnostic database migration tool.

## 📞 Support

If you have any questions or need help, please open an issue on GitHub.

---

**Happy coding! 🎉**
