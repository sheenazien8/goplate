# GoPlate Documentation

Welcome to **GoPlate** - a modern, production-ready Go boilerplate for building REST APIs with best practices built-in.

## What is GoPlate?

GoPlate is a comprehensive Go boilerplate that provides everything you need to build robust REST APIs quickly and efficiently. It combines the power of modern Go frameworks with battle-tested patterns and tools to give you a solid foundation for your next project.

## Key Features

### 🔥 High Performance
- Built on **Fiber** framework - one of the fastest HTTP frameworks for Go
- Optimized for high throughput and low latency
- Efficient memory usage and minimal overhead

### 🗄️ Database Integration
- **GORM** ORM with support for MySQL and PostgreSQL
- Database migrations and seeders
- Connection pooling and optimization
- Automatic model generation

### 🔐 Security First
- JWT authentication middleware
- Request validation with `go-playground/validator`
- CORS support
- Basic authentication for admin endpoints
- Environment-based configuration

### 🛠️ Developer Experience
- Hot reload development server
- Comprehensive CLI tools via Makefile
- Code generation for models, DTOs, and more
- Structured logging with file rotation
- Test coverage reporting

### 📦 Clean Architecture
- Well-organized project structure following Go conventions
- Separation of concerns
- Modular design for easy maintenance
- Scalable codebase organization

### ⏰ Background Processing
- In-memory task queue with worker pools
- CRON-based job scheduling
- Asynchronous task processing
- Configurable worker concurrency

## Quick Overview

```go
// Simple API endpoint example
func (c *Controller) GetUsers(ctx *fiber.Ctx) error {
    users, err := c.userService.GetAll()
    if err != nil {
        return ctx.Status(500).JSON(fiber.Map{
            "error": "Failed to fetch users",
        })
    }
    
    return ctx.JSON(fiber.Map{
        "success": true,
        "data": users,
    })
}
```

## Architecture Highlights

- **MVC Pattern**: Clean separation between Models, Views, and Controllers
- **Middleware Stack**: Authentication, CORS, logging, and error handling
- **Service Layer**: Business logic abstraction
- **Repository Pattern**: Data access abstraction
- **Dependency Injection**: Loose coupling between components

## Use Cases

GoPlate is perfect for:

- **REST APIs**: Build scalable web APIs
- **Microservices**: Create lightweight, focused services  
- **Backend Services**: Power mobile and web applications
- **Data Processing**: Handle background tasks and scheduled jobs
- **Prototyping**: Quickly validate ideas with a solid foundation

## Getting Started

Ready to build something amazing? Let's get you started:

1. **[Quick Start](/quick-start)** - Get up and running in minutes
2. **[Installation](/installation)** - Detailed installation guide
3. **[Configuration](/configuration)** - Configure your environment
4. **[Project Structure](/project-structure)** - Understand the codebase

## Community & Support

- **GitHub**: [sheenazien8/goplate](https://github.com/sheenazien8/goplate)
- **Issues**: Report bugs and request features
- **Discussions**: Ask questions and share ideas

---

**Ready to start building?** Head over to the [Quick Start](/quick-start) guide!