package router

import (
	"heptaber/auth/domain/repository"
	"heptaber/auth/domain/service"
	"heptaber/auth/infrastructure/database"
	"heptaber/auth/infrastructure/web/controllers"

	middleware "heptaber/auth/infrastructure/web/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	var userRepo repository.IUserRepository = repository.NewUserRepository(database.DB)
	var uService service.IUserService = service.NewUserService(userRepo)
	var uController controllers.IUserController = controllers.NewUserController(uService)

	incomingRoutes.Use(middleware.Authenticate())
	userRoutes := incomingRoutes.Group("/api/v1/user")
	{
		userRoutes.GET("/", middleware.HasAdminPermission(), uController.GetUsers())
		userRoutes.GET("/:email", middleware.HasAdminPermission(), uController.GetUserByEmail())
		userRoutes.GET("/me", uController.GetMe())
		userRoutes.PATCH("/role", middleware.HasAdminPermission(), uController.SetUserRole())
		userRoutes.PATCH("/lock", middleware.HasModeratorPermission(), uController.SetUserLock())
	}
}
