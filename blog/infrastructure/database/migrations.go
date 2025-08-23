package database

import (
	"fmt"
	model "heptaber/blog/domain/model"
	"log"
)

// RunMigrations runs SQL migrations and AutoMigrate for blog models.
func RunMigrations() {
    if err := DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error; err != nil {
        log.Fatal(err)
    }

    if err := DB.Exec("DROP TYPE IF EXISTS article_status CASCADE").Error; err != nil {
        log.Fatal(err)
    }
    if err := DB.Exec("CREATE TYPE article_status AS ENUM ('PENDING', 'POSTED', 'REJECTED', 'DRAFT')").Error; err != nil {
        log.Fatal("can not create article_status enum type")
    }

    if err := DB.AutoMigrate(&model.Article{}, &model.Comment{}); err != nil {
        log.Fatal(err)
    }

    if err := DB.Exec("UPDATE articles SET status = 'PENDING' WHERE status IS NULL").Error; err != nil {
        log.Fatal(err)
    }

    fmt.Println("Blog migrations complete.")
}
