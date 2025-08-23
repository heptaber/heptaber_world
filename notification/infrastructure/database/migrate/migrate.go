package main

import (
	"heptaber/notification/infrastructure/database"
)

func init() {
	database.ConnectToDB()
}

func main() {
	database.RunMigrations()
}
