package main

import (
	"os"

	"github.com/sheenazien8/goplate/bootstrap"
	"github.com/sheenazien8/goplate/env"
	"github.com/sheenazien8/goplate/logs"
	"github.com/sheenazien8/goplate/pkg/console"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "console" {
		bootstrap.Init()
		kernel := console.NewKernel()

		if err := kernel.Run(os.Args); err != nil {
			logs.Fatal("Console command failed: ", err.Error())
		}
		return
	}

	app := bootstrap.App()
	port := env.Get("APP_PORT")
	if port == "" {
		port = "8080"
	}

	if err := app.Listen(":" + port); err != nil {
		logs.Fatal("Server won't run: ", err.Error())
	}
}
