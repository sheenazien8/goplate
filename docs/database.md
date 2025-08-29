# Database

Galaplate provides a robust database layer built on top of **GORM** with support for both **MySQL** and **PostgreSQL**. The framework includes automatic migration management, connection pooling, and a comprehensive seeding system.

## Overview

The database system in Galaplate is designed for:
- **Multi-database support**: MySQL and PostgreSQL
- **Migration management**: Version-controlled schema changes via console commands
- **Connection pooling**: Optimized database connections
- **Seeding system**: Populate database with test data
- **Code generation**: Automated model and seeder creation via console commands
- **Developer-friendly CLI**: Powerful console commands for all database operations

## Configuration

Database configuration is managed through environment variables in your `.env` file:

```env
# Database Configuration
DB_CONNECTION=mysql          # or 'postgres'
DB_HOST=localhost
DB_PORT=3306                # 5432 for PostgreSQL
DB_DATABASE=galaplate
DB_USERNAME=root
DB_PASSWORD=password
```

### Supported Drivers

- **MySQL**: Primary support with optimized connection string
- **PostgreSQL**: Full support with SSL configuration options

## Connection Setup

The database connection is established in `db/connect.go`:

```go
// Database connection is initialized globally
var Connect *gorm.DB

// ConnectDB establishes database connection with optimized configuration
func ConnectDB() {
    var gormConfig = &gorm.Config{
        DisableForeignKeyConstraintWhenMigrating: true,
        Logger: logger.New(
            log.New(os.Stdout, "\r\n", log.LstdFlags),
            logger.Config{
                SlowThreshold:             time.Second,
                LogLevel:                  logger.Warn,
                IgnoreRecordNotFoundError: true,
                ParameterizedQueries:      true,
                Colorful:                  true,
            },
        ),
    }
    // Connection logic based on DB_CONNECTION type
}
```

### Connection Features

- **Automatic reconnection**: Handles connection drops gracefully
- **Query logging**: Configurable SQL query logging with slow query detection
- **Parameterized queries**: Protection against SQL injection
- **Connection pooling**: Efficient resource management

## Migrations

Galaplate provides a powerful migration system accessible through console commands, making database schema management simple and efficient.

### Migration Structure

Migrations are stored in `db/migrations/` with the following naming convention:
```
YYYYMMDDHHMMSS_migration_description.sql
```

Example migration file (`20250609004425_create_jobs_table.sql`):

```sql
-- migrate:up
CREATE TABLE jobs (
    id SERIAL PRIMARY KEY,
    type VARCHAR(255) NOT NULL,
    payload JSON,
    state VARCHAR(16) NOT NULL CHECK (state IN ('pending', 'started', 'finished', 'failed')),
    error_msg TEXT,
    attempts INT NULL,
    available_at TIMESTAMP NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    started_at TIMESTAMP NULL,
    finished_at TIMESTAMP NULL
);

-- migrate:down
DROP TABLE IF EXISTS jobs;
```

### Migration Commands

Galaplate provides both modern **Console Commands** and traditional Make commands for database operations:

#### Console Commands

```bash
# Create a new migration
go run main.go console db:create create_users_table

# Run pending migrations
go run main.go console db:up

# Check migration status
go run main.go console db:status

# Rollback last migration
go run main.go console db:down

# Reset database (rollback all migrations and re-run)
go run main.go console db:reset

# Fresh migration (drop tables and run all migrations)
go run main.go console db:fresh

# Run database seeders
go run main.go console db:seed

# List all database commands
go run main.go console list | grep "db:"
```

### Migration Best Practices

- **Always test migrations**: Test both up and down migrations
- **Use transactions**: Ensure atomic migration operations
- **Backup production**: Always backup before running migrations in production
- **Review queries**: Optimize migration queries for large datasets

## Models

Galaplate uses GORM models with struct tags for database mapping. Models are defined in `pkg/models/`.

### Model Example

```go
package models

import (
    "encoding/json"
    "time"
)

type JobState string

const (
    JobPending  JobState = "pending"
    JobStarted  JobState = "started"
    JobFinished JobState = "finished"
    JobFailed   JobState = "failed"
)

type Job struct {
    ID          uint            `gorm:"primaryKey" json:"id"`
    Type        string          `gorm:"not null" json:"type"`
    Payload     json.RawMessage `gorm:"type:text" json:"payload"`
    State       JobState        `gorm:"type:varchar(16);not null" json:"state"`
    ErrorMsg    string          `json:"error_msg"`
    Attempts    int             `json:"attempts"`
    AvailableAt time.Time       `json:"available_at"`
    CreatedAt   time.Time       `json:"created_at"`
    StartedAt   *time.Time      `json:"started_at"`
    FinishedAt  *time.Time      `json:"finished_at"`
}
```

### Model Generation

Generate new models using Console Commands:

```bash
# Generate a new model
go run main.go console make:model User

# Generate a model with specific name
go run main.go console make:model ProductCategory

# List all make commands available
go run main.go console list | grep "make:"
```

**Generated model example:**
```go
package models

import (
    "time"
    "gorm.io/gorm"
)

type User struct {
    ID        uint           `json:"id" gorm:"primaryKey"`
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`

    // Add your fields here
}
```

This creates a new model file with:
- Proper struct definition
- GORM tags for database mapping
- JSON tags for API serialization
- Standard timestamps (CreatedAt, UpdatedAt, DeletedAt)
- Soft delete support

### GORM Features Used

- **Auto-migration**: Automatic table creation and updates
- **Soft deletes**: Records marked as deleted instead of physically removed
- **Hooks**: Before/after save, create, update, delete callbacks
- **Associations**: Relationships between models (HasOne, HasMany, BelongsTo)
- **Scopes**: Reusable query conditions

## Database Seeding

The seeding system allows you to populate your database with test or initial data. Galaplate provides console commands to easily generate and run seeders.

### Seeder Structure

Seeders are located in `db/seeders/` and implement the `Seeder` interface:

```go
func init() {
    registerSeeder("userseeder", &UserSeeder{})
}
```

**Note**: The seeder is automatically registered and will be executed when running `go run main.go console db:seed`.

### Creating Seeders

Generate a new seeder using Console Commands:

```bash
# Create a new seeder
go run main.go console make:seeder UserSeeder

# Create multiple seeders
go run main.go console make:seeder ProductSeeder
go run main.go console make:seeder CategorySeeder
```

This creates a seeder file in `db/seeders/` with the following structure:

```go
package seeders

import (
    "gorm.io/gorm"
    "github.com/sheenazien8/galaplate/pkg/models"
)

type UserSeeder struct{}

func (s *UserSeeder) Seed(db *gorm.DB) error {
    users := []models.User{
        {
            Name:  "Admin User",
            Email: "admin@example.com",
        },
        // Add more users...
    }

    for _, user := range users {
        db.FirstOrCreate(&user, models.User{Email: user.Email})
    }

    return nil
}

func init() {
    registerSeeder("userseeder", &UserSeeder{})
}
```

### Running Seeders

Use Console Commands to run database seeders:

```bash
go run main.go console db:seed

go run main.go console db:fresh  # Drop tables and re-migrate
```

### Seeder Best Practices

- **Idempotent operations**: Use `FirstOrCreate` to avoid duplicates
- **Environment-specific**: Different data for development/staging/production
- **Dependency order**: Seed related data in proper order
- **Error handling**: Graceful handling of seeding failures

## Database Utilities

### Connection Pooling

GORM automatically handles connection pooling with sensible defaults:

```go
// Configure connection pool (optional)
sqlDB, err := db.DB()
sqlDB.SetMaxIdleConns(10)
sqlDB.SetMaxOpenConns(100)
sqlDB.SetConnMaxLifetime(time.Hour)
```

### Query Optimization

- **Preloading**: Load related data efficiently
- **Indexing**: Proper database indexes for common queries
- **Pagination**: Built-in pagination utilities
- **Raw SQL**: Option to use raw SQL for complex queries

```go
// Preloading example
var users []models.User
db.Preload("Profile").Find(&users)

// Pagination example
result := db.Scopes(utils.Paginate(page, pageSize)).Find(&users)
```

### Database Debugging

Enable SQL query logging for development:

```go
// Enable detailed logging
gormConfig.Logger = logger.Default.LogMode(logger.Info)
```

## Security Considerations

- **Parameterized queries**: All queries use parameter binding
- **Connection encryption**: SSL/TLS support for production
- **Access control**: Database user with minimal required permissions
- **Environment isolation**: Separate databases for different environments

## Performance Tips

1. **Use appropriate indexes**: Add indexes for frequently queried fields
2. **Optimize N+1 queries**: Use preloading and joins
3. **Connection pooling**: Configure pool size based on application load
4. **Query analysis**: Monitor slow queries and optimize
5. **Database maintenance**: Regular VACUUM/OPTIMIZE operations
