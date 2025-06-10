package main

import (
	"encoding/json"
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/sheenazien8/goplate/db"
	"github.com/sheenazien8/goplate/env"
	"github.com/sheenazien8/goplate/logs"
	"github.com/sheenazien8/goplate/pkg/queue"
	"github.com/sheenazien8/goplate/pkg/scheduler"
	"github.com/sheenazien8/goplate/pkg/utils"
	"github.com/sheenazien8/goplate/router"
)

func main() {
	screet := env.Get("APP_SCREET")
	if screet == "" {
		logs.Fatal("You must generate the screet key first")
	}

	app := fiber.New(
		fiber.Config{
			ErrorHandler: func(c *fiber.Ctx, err error) error {
				var errResponse utils.GlobalErrorHandlerResp
				errMarshal := json.Unmarshal([]byte(err.Error()), &errResponse)
				if errMarshal != nil {
					logs.Error("Internal Error")
					var e *fiber.Error
					var code int = fiber.StatusInternalServerError
					if errors.As(err, &e) {
						code = e.Code
					}
					logs.Error(err)

					return c.Status(code).JSON(fiber.Map{
						"success": false,
						"message": e.Message,
						"error":   e.Message,
					})
				}

				return c.Status(errResponse.Status).JSON(errResponse)
			},
		},
	)
	db.ConnectDB()
	p := env.Get("APP_PORT")

	router.SetupRouter(app)

	q := queue.New(100)
    q.Start(5)

	sch := scheduler.New()
	sch.RunTasks()
	sch.Start()

	if err := app.Listen(":" + p); err != nil {
		logs.Fatal("Server won't run: ", err.Error())
	}
}
