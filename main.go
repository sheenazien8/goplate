package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"time"

	"github.com/gofiber/fiber/v2"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sheenazien8/goplate/db"
	"github.com/sheenazien8/goplate/env"
	"github.com/sheenazien8/goplate/logs"
	"github.com/sheenazien8/goplate/pkg/utils"
	"github.com/sheenazien8/goplate/router"
)

func main() {
	logPath := "logs/"

	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		if os.Mkdir(logPath, os.ModePerm) != nil {
			logs.Fatal("Error creating log path")
		}
	}

	logFile, err := rotatelogs.New(
		path.Join(logPath, "app_%Y%m%d.log"),
		rotatelogs.WithLinkName(path.Join(logPath, "app.log")),
		rotatelogs.WithRotationTime(24*time.Hour),
		rotatelogs.WithMaxAge(7*24*time.Hour),
	)
	if err != nil {
		fmt.Println("Error creating log file:", err)
	}
	multiLogFile := io.MultiWriter(logFile, os.Stdout)
	logs.SetOutput(multiLogFile)

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

	if err := app.Listen(":" + p); err != nil {
		logs.Fatal("Server won't run: ", err.Error())
	}
}
