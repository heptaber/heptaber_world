package main

import (
	"fmt"
	"log"

	"heptaber/notification/domain/model"
	"heptaber/notification/infrastructure/database"
)

func init() {
	database.ConnectToDB()
}

func main() {
	if err := database.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error; err != nil {
		log.Fatal(err)
	}
	database.DB.AutoMigrate(&model.Notification{})
	fmt.Println("Migration complete.")
}
