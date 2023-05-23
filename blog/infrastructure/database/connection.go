package database

import (
	"heptaber/blog/app/initializers"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	initializers.LoadEnvVariables()
}

var DB *gorm.DB

func ConnectToDB() {
	dsn := os.Getenv("POSTGRES_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database")
	}

	DB = db
}
