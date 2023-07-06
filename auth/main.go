package main

import (
	"os"

	"heptaber/auth/app/cron"
	"heptaber/auth/app/initializers"
	"heptaber/auth/infrastructure/database"
	routes "heptaber/auth/infrastructure/web/router"

	"github.com/gin-contrib/cors"
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
		port = "10010"
	}

	// set up cron job
	cron.SetUpDeleteAllExpiredVerificationCodesJob(database.DB)

	router := gin.New()
	router.Use(gin.Logger())

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.Run(":" + port)
}
