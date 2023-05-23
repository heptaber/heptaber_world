package router

import (
	"heptaber/auth/domain/repository"
	"heptaber/auth/domain/service"
	"heptaber/auth/infrastructure/database"
	"heptaber/auth/infrastructure/web/controllers"
	middleware "heptaber/auth/infrastructure/web/middleware"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	var userRepo repository.IUserRepository = repository.NewUserRepository(database.DB)
	var vCodeRepo repository.IVerificationCodeRepository = repository.NewVerificationCodeRepository(database.DB)
	var tokenRepo repository.ITokenRepository = repository.NewTokenRepository(database.DB)
	var aService service.IAuthService = service.NewAuthService(userRepo, vCodeRepo, tokenRepo)
	var aController controllers.IAuthController = controllers.NewAuthController(aService)

	authRoutes := incomingRoutes.Group("/api/v1/auth")
	{
		authRoutes.POST("/signup", aController.Signup())
		authRoutes.POST("/login", aController.Login())
		authRoutes.POST("/logout", middleware.Authenticate(), aController.Logout())
		authRoutes.POST("/refresh", aController.RefreshToken())
		authRoutes.GET("/verify/:code", aController.VerifyUser())
	}
}
