package db

import (
	"fmt"
	"log"

	"github.com/sheenazien8/goplate/config"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Connect *gorm.DB

func ConnectDB() {
	var err error
	var dsn string
	var dbType = config.Config("DB_CONNECTION")
	var db *gorm.DB

	switch dbType {
	case "postgres":
		p := config.Config("DB_PORT")

		dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			config.Config("DB_HOST"),
			p,
			config.Config("DB_USER"),
			config.Config("DB_PASSWORD"),
			config.Config("DB_NAME"),
		)

		db, err = gorm.Open(
			postgres.Open(dsn),
			&gorm.Config{
				DisableForeignKeyConstraintWhenMigrating: true,
			},
		)

	case "mysql":
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config.Config("DB_USER"),
			config.Config("DB_PASSWORD"),
			config.Config("DB_HOST"),
			config.Config("DB_PORT"),
			config.Config("DB_NAME"),
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
