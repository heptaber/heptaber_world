package database

import (
	"fmt"
	model "heptaber/notification/domain/model"
	"log"
)

// RunMigrations runs SQL migrations and AutoMigrate for notification models.
func RunMigrations() {
    if err := DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error; err != nil {
        log.Fatal(err)
    }

    if err := DB.AutoMigrate(&model.Notification{}); err != nil {
        log.Fatal(err)
    }

    fmt.Println("Notification migrations complete.")
}
