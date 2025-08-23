package database

import (
	"fmt"
	model "heptaber/auth/domain/model"
	"log"
)

// RunMigrations runs SQL migrations and AutoMigrate for auth models.
// It's safe to call multiple times (uses IF NOT EXISTS where appropriate).
func RunMigrations() {
    if err := DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error; err != nil {
        log.Fatal(err)
    }

    // recreate enum type idempotently: drop then create
    if err := DB.Exec("DROP TYPE IF EXISTS user_role CASCADE").Error; err != nil {
        log.Fatal(err)
    }
    if err := DB.Exec("CREATE TYPE user_role AS ENUM ('REGULAR', 'MODERATOR', 'ADMIN')").Error; err != nil {
        log.Fatal("can not create user_role enum type")
    }

    if err := DB.AutoMigrate(&model.User{}, &model.Token{}, &model.VerificationCode{}); err != nil {
        log.Fatal(err)
    }

    if err := DB.Exec("UPDATE users SET role = 'REGULAR' WHERE role IS NULL").Error; err != nil {
        log.Fatal(err)
    }

    fmt.Println("Auth migrations complete.")
}
