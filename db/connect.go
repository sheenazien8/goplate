package db

import (
	"fmt"
	"log"

	"github.com/sheenazien8/goplate/env"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Connect *gorm.DB

func ConnectDB() {
	var err error
	var dsn string
	var dbType = env.Get("DB_CONNECTION")
	var db *gorm.DB

	switch dbType {
	case "postgres":
		p := env.Get("DB_PORT")

		dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			env.Get("DB_HOST"),
			p,
			env.Get("DB_USER"),
			env.Get("DB_PASSWORD"),
			env.Get("DB_NAME"),
		)

		db, err = gorm.Open(
			postgres.Open(dsn),
			&gorm.Config{
				DisableForeignKeyConstraintWhenMigrating: true,
			},
		)

	case "mysql":
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			env.Get("DB_USER"),
			env.Get("DB_PASSWORD"),
			env.Get("DB_HOST"),
			env.Get("DB_PORT"),
			env.Get("DB_NAME"),
		)

		db, err = gorm.Open(
			mysql.Open(dsn),
			&gorm.Config{
				DisableForeignKeyConstraintWhenMigrating: true,
			},
		)
	default:
		log.Panic("Unsupported database type")
	}

	if err != nil {
		log.Panic(err.Error())
	}
	Connect = db

	fmt.Println("Connection Opened to Database")
}
