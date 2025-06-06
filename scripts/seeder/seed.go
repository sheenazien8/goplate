package main

import (
	"os"

	"github.com/sheenazien8/goplate/db"
	"github.com/sheenazien8/goplate/db/seeders"
	"github.com/sheenazien8/goplate/logs"
)

func main() {
	db.ConnectDB()

	selected := ""
	if len(os.Args) > 1 {
		selected = os.Args[1]
	}

	seeder := seeders.NewDatabaseSeeder(selected)

	if err := seeder.Run(db.Connect); err != nil {
		logs.Fatal("Seeding failed:", err)
	}
}
