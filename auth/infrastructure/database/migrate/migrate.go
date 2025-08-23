package main

import (
	"heptaber/auth/infrastructure/database"
)

func init() {
	database.ConnectToDB()
}

func main() {
	database.RunMigrations()
}
