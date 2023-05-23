package database

import (
	model "heptaber/blog/domain/model"
	"log"
)

func SyncDatabase() {
	if err := DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error; err != nil {
		log.Fatal(err)
	}
	DB.AutoMigrate(&model.Article{}, &model.Comment{})
}
