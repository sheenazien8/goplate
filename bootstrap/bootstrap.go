package bootstrap

import (
	"encoding/json"
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/sheenazien8/goplate/db"
	"github.com/sheenazien8/goplate/env"
	"github.com/sheenazien8/goplate/logs"
	"github.com/sheenazien8/goplate/pkg/queue"
	"github.com/sheenazien8/goplate/pkg/scheduler"
	"github.com/sheenazien8/goplate/pkg/utils"
	"github.com/sheenazien8/goplate/router"
)

func App() *fiber.App {
	screet := env.Get("APP_SCREET")
	if screet == "" {
		logs.Fatal("You must generate the screet key first")
	}

	engine := html.New("./templates", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			var errResponse utils.GlobalErrorHandlerResp
			if json.Unmarshal([]byte(err.Error()), &errResponse) != nil {
				var e *fiber.Error
				code := fiber.StatusInternalServerError
				message := "Internal Server Error"
				if errors.As(err, &e) {
					code = e.Code
					message = e.Message
				}
				logs.Error(err)

				return c.Status(code).JSON(fiber.Map{
					"success": false,
					"message": message,
					"error":   message,
				})
			}
			return c.Status(errResponse.Status).JSON(errResponse)
		},
	})

	// Connect DB (can be swapped with test DB)
	db.ConnectDB()

	// Setup routes
	router.SetupRouter(app)

	// Start background jobs (can be skipped in tests)
	q := queue.New(100)
	q.Start(5)

	sch := scheduler.New()
	sch.RunTasks()
	sch.Start()

	return app
}

func Init() {
	screet := env.Get("APP_SCREET")
	if screet == "" {
		logs.Fatal("You must generate the screet key first")
	}

	// Connect DB for console commands
	db.ConnectDB()
}
