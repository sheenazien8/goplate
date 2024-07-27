package main

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sheenazien8/goplate/config"
	"github.com/sheenazien8/goplate/db"
	"github.com/sheenazien8/goplate/pkg/utils"
	"github.com/sheenazien8/goplate/router"
)

func main() {
	screet := config.Config("APP_SCREET")
	if screet == "" {
		log.Panic("You must generate the screet key first")
	}

	app := fiber.New(
		fiber.Config{
			ErrorHandler: func(c *fiber.Ctx, err error) error {
				var errResponse utils.GlobalErrorHandlerResp
				errMarshal := json.Unmarshal([]byte(err.Error()), &errResponse)
				if errMarshal != nil {
					log.Println("Internal Error")
					var e *fiber.Error
					var code int = fiber.StatusInternalServerError
					if errors.As(err, &e) {
						code = e.Code
					}

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
	p := config.Config("APP_PORT")

	router.SetupRouter(app)

	if err := app.Listen(":" + p); err != nil {
		log.Panic("Server won't run", err.Error())
	}
}
