# Console Commands

GoPlate includes a powerful console command system that provides code generation, database management, and custom command capabilities. All console commands are executed through the main application entry point.

## Basic Usage

```bash
# General syntax
go run main.go console <command> [arguments]

# List all available commands
go run main.go console list

# Get help
go run main.go console
```

## Available Commands

### Database Commands

#### `db:create`
Create a new database migration file.

```bash
go run main.go console db:create create_users_table
```

#### `db:up`
Run pending database migrations.

```bash
go run main.go console db:up
```

#### `db:down`
Rollback the last database migration.

```bash
go run main.go console db:down
```

#### `db:status`
Show current migration status.

```bash
go run main.go console db:status
```

#### `db:fresh`
Drop all tables and re-run all migrations.

```bash
go run main.go console db:fresh
```

#### `db:reset`
Rollback all migrations and re-run them.

```bash
go run main.go console db:reset
```

#### `db:seed`
Run database seeders.

```bash
go run main.go console db:seed
```

### Code Generation Commands

#### `make:model`
Generate a new model file.

```bash
go run main.go console make:model User
```

**Generated file:** `pkg/models/user.go`

**Example output:**
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
}
```

#### `make:dto`
Generate a new Data Transfer Object (DTO).

```bash
go run main.go console make:dto UserDto
```

**Generated file:** `pkg/dto/user_dto.go`

#### `make:job`
Generate a new background job.

```bash
go run main.go console make:job ProcessEmailJob
```

**Generated file:** `pkg/queue/jobs/process_email_job.go`

#### `make:seeder`
Generate a new database seeder.

```bash
go run main.go console make:seeder UserSeeder
```

**Generated file:** `db/seeders/user_seeder.go`

#### `make:cron`
Generate a new cron job for the scheduler.

```bash
go run main.go console make:cron DailyReportCron
```

**Generated file:** `pkg/scheduler/daily_report_cron.go`

### Utility Commands

#### `list`
Display all available console commands with descriptions.

```bash
go run main.go console list
```

#### `example`
Run an example command to test the console system.

```bash
go run main.go console example
```

#### `interactive`
Start an interactive demonstration of console features.

```bash
go run main.go console interactive
```

## Creating Custom Commands

### Step 1: Create Command File

Create a new command file in `pkg/console/commands/`:

```go
// pkg/console/commands/my_custom_command.go
package commands

import (
    "fmt"
)

type MyCustomCommand struct{}

func (c *MyCustomCommand) GetSignature() string {
    return "my:custom"
}

func (c *MyCustomCommand) GetDescription() string {
    return "My custom command description"
}

func (c *MyCustomCommand) Execute(args []string) error {
    fmt.Println("Executing my custom command!")

    if len(args) > 0 {
        fmt.Printf("Arguments: %v\n", args)
    }

    return nil
}
```

### Step 2: Register Command

Add your command to the registration in `pkg/console/commands.go`:

```go
func (k *Kernel) RegisterCommands() {
    // Example command (you can remove this)
    k.Register(&commands.ExampleCommand{})

    // Interactive demo command
    k.Register(&commands.InteractiveCommand{})

    // Register your custom command
    k.Register(&commands.MyCustomCommand{})
}
```

### Step 3: Use Your Command

```bash
go run main.go console my:custom arg1 arg2
```

## Command Interface

All commands must implement the `Command` interface:

```go
type Command interface {
    GetSignature() string    // Command name (e.g., "make:model")
    GetDescription() string  // Brief description for help
    Execute(args []string) error // Command logic
}
```

### Command Naming Conventions

- Use `:` to separate command namespaces (e.g., `make:model`, `db:up`)
- Use lowercase with hyphens for multi-word commands (e.g., `cache:clear`)
- Group related commands under the same namespace

### Best Practices

1. **Error Handling**: Always return meaningful errors from `Execute()`
2. **Argument Validation**: Validate required arguments early
3. **User Feedback**: Provide clear success/failure messages
4. **Help Text**: Include usage examples in descriptions

### Example: Advanced Custom Command

```go
type DeployCommand struct{}

func (c *DeployCommand) GetSignature() string {
    return "deploy"
}

func (c *DeployCommand) GetDescription() string {
    return "Deploy application to specified environment"
}

func (c *DeployCommand) Execute(args []string) error {
    if len(args) < 1 {
        return fmt.Errorf("environment required. Usage: deploy <environment>")
    }

    environment := args[0]

    fmt.Printf("Deploying to %s environment...\n", environment)

    // Your deployment logic here

    fmt.Printf("Successfully deployed to %s!\n", environment)
    return nil
}
```

## Integration with Make Commands

Console commands work alongside traditional Make commands:

```bash
# These are equivalent:
go run main.go console db:up

# These are equivalent:
go run main.go console make:model User
# (No direct make equivalent for code generation)
```

## Environment Integration

Console commands automatically load your `.env` configuration:

```go
func (c *MyCommand) Execute(args []string) error {
    // Environment variables are available
    dbHost := os.Getenv("DB_HOST")
    appName := os.Getenv("APP_NAME")

    // Your command logic
    return nil
}
```

## Debugging Commands

Enable verbose output for debugging:

```bash
# Set debug mode in .env
APP_DEBUG=true

# Run your command
go run main.go console my:command
```

## Performance Tips

1. **Lazy Loading**: Only import packages when needed
2. **Early Exit**: Validate inputs before heavy operations
3. **Progress Feedback**: Show progress for long-running commands
4. **Resource Cleanup**: Always clean up resources in commands

---

**Next Steps:**
- **[Background Tasks](/background-tasks)** - Learn about job processing
- **[Database](/database)** - Understand database operations
- **[API Reference](/api-reference)** - Explore API endpoints
