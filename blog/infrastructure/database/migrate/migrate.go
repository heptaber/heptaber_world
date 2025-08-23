package main

import (
	"heptaber/blog/infrastructure/database"
)

func init() {
	database.ConnectToDB()
}

func main() {
	database.RunMigrations()
}
