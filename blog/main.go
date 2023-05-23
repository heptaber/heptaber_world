package main

import (
	"os"

	"heptaber/blog/app/initializers"
	"heptaber/blog/infrastructure/database"

	routes "heptaber/blog/infrastructure/web/router"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	database.ConnectToDB()
	database.SyncDatabase()
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "10030"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.ArticleRoute(router)

	router.Run(":" + port)
}
