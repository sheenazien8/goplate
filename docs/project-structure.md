# Project Structure

GoPlate follows Go best practices and conventions for project organization. This guide explains the purpose of each directory and file in the project.

## Overview

```
goplate/
├── db/                     # Database related files
│   ├── migrations/        # SQL migration files
│   ├── seeders/          # Database seeders
│   └── connect.go        # Database connection setup
├── docs/                  # Documentation (Docsify)
│   ├── index.html        # Docsify configuration
│   ├── README.md         # Main documentation
│   └── ...              # Other documentation files
├── env/                   # Environment configuration
│   └── config.go         # Environment variable utilities
├── logs/                  # Logging utilities
│   └── logger.go         # Logging configuration
├── middleware/            # HTTP middleware
│   └── auth.go           # Authentication middleware
├── pkg/                   # Main application packages
│   ├── controllers/      # HTTP request handlers
│   ├── models/          # Database models
│   ├── queue/           # Background job queue
│   ├── scheduler/       # CRON job scheduler
│   └── utils/           # Utility functions
├── router/                # Route definitions
│   └── router.go        # HTTP routes setup
├── scripts/              # Helper scripts
│   ├── seeder/          # Database seeding scripts
│   ├── stubs/           # Code generation templates
│   └── *.sh             # Shell scripts for development
├── storage/              # Application storage
│   └── logs/            # Log files
├── views/                # HTML templates
│   └── logs.html        # Log viewer template
├── .env.example          # Environment variables template
├── .gitignore           # Git ignore rules
├── go.mod               # Go module definition
├── go.sum               # Go module checksums
├── main.go              # Application entry point
├── Makefile             # Development commands
└── README.md            # Project documentation
```

## Directory Details

### `/db` - Database Layer

All database-related functionality.

```
db/
├── migrations/           # SQL migration files
│   └── 20250609004425_create_jobs_table.sql
├── seeders/             # Database seeders
│   └── main.go
└── connect.go           # Database connection
```

**Key Files:**
- `connect.go` - Database connection setup and configuration
- `migrations/` - SQL files for database schema changes
- `seeders/` - Go files for populating test data

**Example Migration:**
```sql
-- 20250609004425_create_jobs_table.sql
CREATE TABLE jobs (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    type VARCHAR(255) NOT NULL,
    payload TEXT,
    state VARCHAR(16) NOT NULL,
    -- ... other fields
);
```

### `/env` - Environment Configuration

Environment variable management and configuration utilities.

```
env/
└── config.go            # Environment utilities
```

**Purpose:**
- Load environment variables from `.env` file
- Provide type-safe access to configuration
- Default value handling

**Example:**
```go
// env/config.go
func Get(key string) string {
    return os.Getenv(key)
}

func GetInt(key string, defaultValue int) int {
    // Implementation
}
```

### `/logs` - Logging System

Centralized logging configuration and utilities.

```
logs/
└── logger.go            # Logging setup
```

**Features:**
- Structured JSON logging
- File rotation (daily)
- Multiple log levels
- Contextual logging

**Example:**
```go
// logs/logger.go
func Info(args ...interface{}) {
    logger.Info(args...)
}

func Error(args ...interface{}) {
    logger.Error(args...)
}
```

### `/middleware` - HTTP Middleware

Reusable HTTP middleware components.

```
middleware/
└── auth.go              # Authentication middleware
```

**Common Middleware:**
- Authentication (JWT, Basic Auth)
- CORS handling
- Request logging
- Rate limiting
- Error handling

**Example:**
```go
// middleware/auth.go
func BasicAuth() fiber.Handler {
    return basicauth.New(basicauth.Config{
        Users: map[string]string{
            username: password,
        },
    })
}
```

### `/pkg` - Core Application Logic

Main application packages following Go conventions.

```
pkg/
├── controllers/         # HTTP request handlers
│   └── log_controller.go
├── models/             # Database models
│   └── job.go
├── queue/              # Background job system
│   ├── jobs/          # Job implementations
│   └── queue.go       # Queue management
├── scheduler/          # CRON scheduler
│   └── main.go
└── utils/              # Utility functions
    ├── bcrypt.go      # Password hashing
    ├── pagination.go  # Pagination helpers
    └── validator.go   # Request validation
```

#### Controllers
Handle HTTP requests and responses.

```go
// pkg/controllers/log_controller.go
type LogController struct {}

func (c *LogController) ShowLogsPage(ctx *fiber.Ctx) error {
    // Handle log viewer request
}
```

#### Models
Define database entities and business logic.

```go
// pkg/models/job.go
type Job struct {
    ID          uint            `gorm:"primaryKey" json:"id"`
    Type        string          `gorm:"not null" json:"type"`
    Payload     json.RawMessage `gorm:"type:text" json:"payload"`
    State       JobState        `gorm:"type:varchar(16);not null" json:"state"`
    // ... other fields
}
```

#### Queue System
Background job processing with worker pools.

```go
// pkg/queue/queue.go
type Queue struct {
    tasks chan Task
    wg    sync.WaitGroup
}

func (q *Queue) Start(workerCount int) {
    // Start worker goroutines
}
```

#### Utilities
Common helper functions and utilities.

```go
// pkg/utils/validator.go
func ValidateStruct(s interface{}) error {
    // Validation logic
}
```

### `/router` - Route Definitions

HTTP route configuration and setup.

```
router/
└── router.go            # Route definitions
```

**Purpose:**
- Define all HTTP routes
- Apply middleware to routes
- Group related routes

**Example:**
```go
// router/router.go
func SetupRouter(app *fiber.App) {
    app.Use(cors.New())

    app.Get("/", healthCheck)
    app.Get("/logs", middleware.BasicAuth(), logController.ShowLogsPage)

    // API routes
    api := app.Group("/api")
    api.Get("/users", userController.GetUsers)
}
```

### `/scripts` - Development Scripts

Helper scripts for development and deployment.

```
scripts/
├── seeder/              # Database seeding
├── stubs/              # Code generation templates
├── tinker/             # Interactive shell
├── generate_model.sh   # Model generation
├── generate_dto.sh     # DTO generation
├── migrate.sh          # Database migrations
└── ...                 # Other utility scripts
```

**Key Scripts:**
- `migrate.sh` - Database migration management
- `generate_model.sh` - Generate new models
- `generate_dto.sh` - Generate DTOs
- `generate_job.sh` - Generate background jobs

### `/storage` - Application Storage

Runtime storage for logs, uploads, and temporary files.

```
storage/
└── logs/               # Application log files
    └── app.2025-06-24.log
```

**Purpose:**
- Store application logs
- Temporary file storage
- Upload storage (when implemented)

### `/views` - HTML Templates

HTML templates for web interfaces.

```
views/
└── logs.html           # Log viewer template
```

**Purpose:**
- Admin interfaces
- Email templates
- Web dashboards

## Configuration Files

### `main.go` - Application Entry Point

The main application file that:
- Initializes the Fiber app
- Sets up middleware
- Connects to database
- Starts background services
- Configures error handling

```go
func main() {
    // Load environment
    screet := env.Get("APP_SCREET")

    // Create Fiber app
    app := fiber.New(fiber.Config{
        ErrorHandler: globalErrorHandler,
        Views: engine,
    })

    // Setup database
    db.ConnectDB()

    // Setup routes
    router.SetupRouter(app)

    // Start background services
    queue := queue.New(100)
    queue.Start(5)

    scheduler := scheduler.New()
    scheduler.Start()

    // Start server
    app.Listen(":" + env.Get("APP_PORT"))
}
```

### `Makefile` - Development Commands

Provides convenient commands for development:

```makefile
# Development
dev:                    # Start with hot reload
run:                    # Build and run
build:                  # Build binary
test:                   # Run tests

# Database
db-up:                  # Run migrations
db-down:                # Rollback migration
db-fresh:               # Reset and migrate

# Code Generation
model:                  # Generate model
dto:                    # Generate DTO
```

### `.env.example` - Environment Template

Template for environment configuration:

```env
APP_NAME=GoPlate
APP_ENV=local
APP_DEBUG=true
APP_PORT=8080
APP_SCREET=your-secret-key

DB_CONNECTION=mysql
DB_HOST=localhost
DB_PORT=3306
DB_DATABASE=goplate
DB_USERNAME=root
DB_PASSWORD=
```

## Design Patterns

### MVC Architecture

GoPlate follows the Model-View-Controller pattern:

- **Models** (`/pkg/models/`) - Data layer and business logic
- **Views** (`/views/`) - Presentation layer (HTML templates)
- **Controllers** (`/pkg/controllers/`) - Request handling logic

### Repository Pattern

Database access is abstracted through repositories:

```go
type UserRepository interface {
    GetByID(id uint) (*User, error)
    Create(user *User) error
    Update(user *User) error
    Delete(id uint) error
}
```

### Service Layer

Business logic is encapsulated in service layers:

```go
type UserService struct {
    repo UserRepository
}

func (s *UserService) CreateUser(data CreateUserDTO) (*User, error) {
    // Business logic
}
```

### Dependency Injection

Dependencies are injected through constructors:

```go
func NewUserController(service UserService) *UserController {
    return &UserController{
        service: service,
    }
}
```

## Best Practices

### File Naming

- Use snake_case for files: `user_controller.go`
- Use descriptive names: `email_service.go`
- Group related files in packages

### Package Organization

- Keep packages focused and cohesive
- Avoid circular dependencies
- Use interfaces for abstraction

### Error Handling

- Return errors explicitly
- Use structured error types
- Log errors with context

### Testing

- Place tests next to source files: `user_test.go`
- Use table-driven tests
- Mock external dependencies

## Extending the Structure

### Adding New Features

1. **Create model** in `/pkg/models/`
2. **Add controller** in `/pkg/controllers/`
3. **Define routes** in `/router/router.go`
4. **Add migrations** in `/db/migrations/`
5. **Write tests** alongside source files

### Adding Middleware

1. Create middleware in `/middleware/`
2. Apply in `/router/router.go`
3. Document usage and configuration

### Adding Background Jobs

1. Implement job in `/pkg/queue/jobs/`
2. Register job in queue system
3. Dispatch jobs from controllers

---

This structure provides a solid foundation for building scalable Go applications while maintaining clean code organization and following Go best practices.
