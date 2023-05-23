package router

import (
	"heptaber/blog/domain/repository"
	"heptaber/blog/domain/service"
	"heptaber/blog/infrastructure/database"

	"heptaber/blog/infrastructure/web/controllers"
	middleware "heptaber/blog/infrastructure/web/middleware"

	"github.com/gin-gonic/gin"
)

func ArticleRoute(incomingRoutes *gin.Engine) {
	var articleRepo repository.IArticleRepository = repository.NewArticleRepository(database.DB)
	var commentRepo repository.ICommentRepository = repository.NewCommentRepository(database.DB)
	var cService service.ICommentService = service.NewCommentService(commentRepo)
	var aService service.IArticleService = service.NewArticleService(articleRepo)
	var aController controllers.IArticleController = controllers.NewArticleController(aService, cService)

	incomingRoutes.Use(middleware.Authenticate())
	articleRoutes := incomingRoutes.Group("/api/v1/articles")
	{
		articleRoutes.GET("", aController.GetArticles()) // get all posted by page
		articleRoutes.POST("", aController.CreateArticle())
		articleRoutes.GET("/all/status/:status", middleware.HasModeratorPermission(), aController.GetArticlesByStatus())
		articleRoutes.GET("/my/status/:status", aController.GetMyArticlesByStatus())
		articleRoutes.PUT("/status", middleware.HasModeratorPermission(), aController.SetArticleStatus())
		articleRoutes.GET("/article/:article_id", aController.GetArticle())
		articleRoutes.GET("/article/comments/:article_id", aController.GetArticleComments())
		articleRoutes.POST("/article/comments", aController.CreateComment())
		articleRoutes.PUT("/article/comments/visibility", middleware.HasModeratorPermission(), aController.SetCommentIsVisible())
	}
}
