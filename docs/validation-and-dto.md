# DTOs & Validation

## Overview

GoPlate uses Data Transfer Objects (DTOs) to define the structure of request and response payloads, and leverages `go-playground/validator` for robust request validation. This ensures type safety, clear API contracts, and reliable error handling.

- **DTOs**: Structs for input/output data
- **Validation**: Tag-based, automatic, and extensible
- **Error Handling**: Consistent validation error responses

## Defining DTOs

DTOs are Go structs, typically placed in `pkg/dto/` or alongside controllers.

```go
// pkg/dto/user_dto.go
type CreateUserDTO struct {
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required,min=8"`
}
```

## Using DTOs in Controllers

Bind and validate DTOs in your handler:

```go
func (c *UserController) Register(ctx *fiber.Ctx) error {
    var dto CreateUserDTO
    if err := ctx.BodyParser(&dto); err != nil {
        return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
    }
    if err := validator.New().Struct(dto); err != nil {
        return ctx.Status(422).JSON(fiber.Map{"error": "Validation failed", "details": err.Error()})
    }
    // Proceed with valid data
}
```

## Validation Tags

- `required` — Field must be present
- `email` — Must be a valid email
- `min`, `max` — Length or value constraints
- `len` — Exact length
- `oneof` — Must match one of the listed values

**Example:**
```go
type ProductDTO struct {
    Name  string  `json:"name" validate:"required,min=3"`
    Price float64 `json:"price" validate:"required,gt=0"`
}
```

## Custom Validation

You can register custom validation functions:

```go
validate := validator.New()
validate.RegisterValidation("username", func(fl validator.FieldLevel) bool {
    return len(fl.Field().String()) >= 3
})
```

## Validation Error Response

Validation errors return a 422 status and details:

```json
{
  "error": "Validation failed",
  "details": {
    "email": ["email is required", "email must be valid"],
    "password": ["password is required", "password must be at least 8 characters"]
  }
}
```

## Best Practices

- Define DTOs for all request/response bodies
- Use validation tags for all fields
- Return clear error messages for invalid input
- Keep DTOs and models separate

---

**Next:** See [API Reference](/api-reference) for endpoint details.