package main

import (
	"fmt"
	"log"

	"heptaber/blog/domain/model"
	"heptaber/blog/infrastructure/database"
)

func init() {
	database.ConnectToDB()
}

func main() {
	if err := database.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error; err != nil {
		log.Fatal(err)
	}
	if err := database.DB.Exec("DROP TYPE IF EXISTS article_status CASCADE").Error; err != nil {
		log.Fatal(err)
	}
	if err := database.DB.Exec("CREATE TYPE article_status AS ENUM ('PENDING', 'POSTED', 'REJECTED', 'DRAFT')").Error; err != nil {
		log.Fatal("can not create user_role enum type")
	}
	database.DB.AutoMigrate(&model.Article{}, &model.Comment{})
	if err := database.DB.Exec("UPDATE articles SET status = 'PENDING' WHERE status IS NULL").Error; err != nil {
		log.Fatal(err)
	}
	fmt.Println("Migration complete.")
}
