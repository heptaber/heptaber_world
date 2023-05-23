package database

import (
	"log"
	"os"

	"heptaber/notification/app/initializers"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	initializers.LoadEnvVariables()
	dsn := os.Getenv("POSTGRES_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database")
	}

	DB = db
}
