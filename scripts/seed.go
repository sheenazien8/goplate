package main

import (
	"github.com/sheenazien8/goplate/db"
	"github.com/sheenazien8/goplate/db/seeders"
	"github.com/sheenazien8/goplate/logs"
)


func main() {
	db.ConnectDB()

	seeder := seeders.NewDatabaseSeeder()

	if err := seeder.Run(db.Connect); err != nil {
		logs.Fatal("Seeding failed:", err)
	}
}
