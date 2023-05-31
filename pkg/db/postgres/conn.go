package postgres

import (
	"fmt"
	"os"

	"github.com/fajarabdillahfn/todo-grpc/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenDB() *gorm.DB {
	host := os.Getenv("SQL_HOST")
	port := os.Getenv("SQL_PORT")
	user := os.Getenv("SQL_USER")
	password := os.Getenv("SQL_PASS")
	dbname := os.Getenv("SQL_DB")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	openConnection := postgres.Open(psqlInfo)

	DB, err := gorm.Open(openConnection, &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	if os.Getenv("SQL_MIGRATE") == "true" {
		DB = DB.Set("gorm:auto_preload", true)

		if err := autoMigrate(DB); err != nil {
			panic("failed to migrate, caused: " + err.Error())
		}
	}

	return DB
}

func autoMigrate(DB *gorm.DB) error {
	err := DB.AutoMigrate(
		model.Task{},
	)

	if err != nil {
		return err
	}

	return nil
}
