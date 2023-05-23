package main

import (
	"fmt"
	"log"

	"heptaber/auth/domain/model"
	"heptaber/auth/infrastructure/database"
)

func init() {
	database.ConnectToDB()
}

func main() {
	if err := database.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error; err != nil {
		log.Fatal(err)
	}
	if err := database.DB.Exec("DROP TYPE IF EXISTS user_role CASCADE").Error; err != nil {
		log.Fatal(err)
	}
	if err := database.DB.Exec("CREATE TYPE user_role AS ENUM ('REGULAR', 'MODERATOR', 'ADMIN')").Error; err != nil {
		log.Fatal("can not create user_role enum type")
	}
	database.DB.AutoMigrate(&model.User{}, &model.Token{}, &model.VerificationCode{})
	if err := database.DB.Exec("UPDATE users SET role = 'REGULAR' WHERE role IS NULL").Error; err != nil {
		log.Fatal(err)
	}
	fmt.Println("Migration complete.")
}
