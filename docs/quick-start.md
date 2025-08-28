# Quick Start

Get up and running with GoPlate in just a few minutes!

## Prerequisites

Before you begin, make sure you have:

- **Go 1.22+** installed on your system
- **MySQL** or **PostgreSQL** database server
- **Git** for version control
- **Make** utility (usually pre-installed on Unix systems)

## Installation Methods

### Method 1: Using the Install Script (Recommended)

The fastest way to get started:

```bash
curl -sSL https://raw.githubusercontent.com/sheenazien8/goplate/master/install.sh | bash
```

### Method 2: Manual Installation

```bash
# Install the CLI tool
go install github.com/sheenazien8/goplate/cmd/goplate@latest

# Or clone the repository directly
git clone https://github.com/sheenazien8/goplate.git
cd goplate
```

## Create Your First Project

### Using GoPlate CLI

```bash
# Create a new project
goplate my-awesome-api
cd my-awesome-api

# Install dependencies
go mod tidy
```

### Using Git Clone

```bash
# Clone the repository
git clone https://github.com/sheenazien8/goplate.git my-awesome-api
cd my-awesome-api

# Remove git history and initialize your own
rm -rf .git
git init
git add .
git commit -m "Initial commit"

# Install dependencies
go mod tidy
```

## Environment Configuration

1. **Copy the environment template:**
   ```bash
   cp .env.example .env
   ```

2. **Edit your `.env` file:**
   ```env
   APP_NAME=MyAwesomeAPI
   APP_ENV=local
   APP_DEBUG=true
   APP_URL=http://localhost
   APP_PORT=8080
   APP_SCREET=your-super-secret-key-here

   # Database Configuration
   DB_CONNECTION=mysql
   DB_HOST=localhost
   DB_PORT=3306
   DB_DATABASE=my_awesome_api
   DB_USERNAME=root
   DB_PASSWORD=your_password

   # Basic Auth for Admin Endpoints
   BASIC_AUTH_USERNAME=admin
   BASIC_AUTH_PASSWORD=secure_password
   ```

3. **Generate a secure secret key:**
   ```bash
   # Generate a random 32-character secret
   openssl rand -base64 32
   ```

## Database Setup

### Create Database

<!-- tabs:start -->

#### **MySQL**

```sql
CREATE DATABASE my_awesome_api CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

#### **PostgreSQL**

```sql
CREATE DATABASE my_awesome_api;
```

<!-- tabs:end -->

### Run Migrations

```bash
# View all available console commands
go run main.go console list

# Create database (if needed)
go run main.go console db:create

# Run database migrations
go run main.go console db:up

# Check migration status
go run main.go console db:status
```

## Start Development Server

### Option 1: Hot Reload Development (Recommended)

```bash
# Start development server with hot reload
make dev
```

This will:
- Install `reflex` if not already installed
- Watch for file changes
- Automatically rebuild and restart the server

### Option 2: Manual Build and Run

```bash
# Build and run
make run

# Or build separately
make build
./server
```

## Verify Installation

Once your server is running, you can test it:

### 1. Basic Health Check

```bash
curl http://localhost:8080/
```

**Expected Response:**
```
Hello world
```

### 2. Access Logs (Admin Endpoint)

```bash
curl -u admin:secure_password http://localhost:8080/logs
```

This should return an HTML page showing your application logs.

### 3. Check Server Logs

Your application logs will be available in:
```
storage/logs/app.YYYY-MM-DD.log
```

## Next Steps

Congratulations! 🎉 Your GoPlate application is now running. Here's what you can do next:

### 1. Explore the Codebase
- **[Project Structure](/project-structure)** - Understand how the code is organized
- **[Configuration](/configuration)** - Learn about all configuration options

### 2. Build Your API
- **[Models & DTOs](/models-dtos)** - Create data models
- **[Controllers](/controllers)** - Build API endpoints
- **[Database](/database)** - Work with databases

### 3. Add Features
- **[Authentication](/authentication)** - Implement user authentication
- **[Background Tasks](/background-tasks)** - Process tasks asynchronously
- **[Validation](/validation)** - Validate incoming requests

### 4. Development Tools
- **[Code Generation](/code-generation)** - Generate boilerplate code
- **[Testing](/testing)** - Write and run tests
- **[Migrations](/migrations)** - Manage database schema

## Common Issues

### Port Already in Use

If port 8080 is already in use:

```bash
# Change the port in your .env file
APP_PORT=3000

# Or set it temporarily
APP_PORT=3000 make run
```

### Database Connection Issues

1. **Check your database is running:**
   ```bash
   # MySQL
   brew services start mysql
   # or
   sudo systemctl start mysql

   # PostgreSQL
   brew services start postgresql
   # or
   sudo systemctl start postgresql
   ```

2. **Verify connection details in `.env`**

3. **Test database connection:**
   ```bash
   make db-connect
   ```

### Missing Dependencies

If you encounter missing dependencies:

```bash
# Install all development dependencies
make install-deps

# Tidy go modules
make tidy
```

## Development Commands

Here are the most commonly used commands during development:

```bash
# Development
make dev              # Start with hot reload
make run              # Build and run once
make build            # Build binary only
make clean            # Clean build artifacts

# Database (Console Commands)
go run main.go console db:up        # Run migrations
go run main.go console db:down      # Rollback migration
go run main.go console db:status    # Check migration status
go run main.go console db:fresh     # Reset and migrate
go run main.go console db:seed      # Run database seeders

# Database (Make Commands - Alternative)
make db-up            # Run migrations
make db-down          # Rollback migration
make db-status        # Check migration status
make db-fresh         # Reset and migrate

# Code Quality
make fmt              # Format code
make test             # Run tests
make test-coverage    # Run tests with coverage

# Code Generation (Console Commands)
go run main.go console make:model User      # Generate new model
go run main.go console make:dto UserDto     # Generate new DTO
go run main.go console make:job ProcessData # Generate background job
go run main.go console make:seeder UserSeeder # Generate database seeder

# Console System
go run main.go console list         # List all available commands
go run main.go console example      # Run example command
go run main.go console interactive  # Interactive demo
```

## Getting Help

- **[Full Documentation](/)**
- **[API Reference](/api-reference)**
- **[Examples](/examples)**
- **[GitHub Issues](https://github.com/sheenazien8/goplate/issues)**

---

**You're all set!** 🚀 Start building your amazing API with GoPlate!
