package database

import (
	"heptaber/blog/app/initializers"
	"log"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectToDB connects to Postgres with retries and exponential backoff.
// Retry behavior can be configured with env vars:
// POSTGRES_MAX_RETRIES (default 10) and POSTGRES_RETRY_WAIT_SECONDS (default 1)
func ConnectToDB() {
	initializers.LoadEnvVariables()
	dsn := os.Getenv("POSTGRES_DSN")

	maxRetries := 10
	if v := os.Getenv("POSTGRES_MAX_RETRIES"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n > 0 {
			maxRetries = n
		}
	}

	waitSeconds := 1
	if v := os.Getenv("POSTGRES_RETRY_WAIT_SECONDS"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n > 0 {
			waitSeconds = n
		}
	}

	wait := time.Duration(waitSeconds) * time.Second

	var err error
	for attempt := 1; attempt <= maxRetries; attempt++ {
		var db *gorm.DB
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			DB = db
			log.Printf("connected to database on attempt %d/%d", attempt, maxRetries)
			return
		}

		if attempt == 1 || attempt%3 == 0 || attempt == maxRetries {
			log.Printf("[warning] failed to connect to database (attempt %d/%d): %v", attempt, maxRetries, err)
		}

		if attempt == maxRetries {
			break
		}

		time.Sleep(wait)
		wait = min(wait * 2, 30 * time.Second)
	}

	log.Fatal("failed to connect to database after retries: ", err)
}
