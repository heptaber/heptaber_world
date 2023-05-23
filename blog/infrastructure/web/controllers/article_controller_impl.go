package controllers

import (
	"heptaber/blog/app/helper"
	"heptaber/blog/domain/model"
	"heptaber/blog/domain/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type articleController struct {
	articleService service.IArticleService
	commentService service.ICommentService
}

func NewArticleController(articleService service.IArticleService, commentService service.ICommentService,
) *articleController {
	return &articleController{
		articleService: articleService,
		commentService: commentService,
	}
}

func (ac *articleController) GetArticles() gin.HandlerFunc {
	return func(c *gin.Context) {
		recordPerPage, err := strconv.Atoi(c.DefaultQuery("recordPerPage", "20"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid page size param"})
			return
		}
		page, err := strconv.Atoi(c.DefaultQuery("pageNumber", "1"))
		if err != nil || page < 1 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid page number param"})
			return
		}

		articles, totalCount, err := ac.articleService.GetArticlesByStatus(recordPerPage, page, model.POSTED)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while retrieving articles"})
			return
		}

		shortArticlesDTO := ConvertArticlesToShortArticlesDTO(*articles)

		c.JSON(http.StatusOK, gin.H{
			"articles":   shortArticlesDTO,
			"totalCount": totalCount,
		})
	}
}

func (ac *articleController) CreateArticle() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := helper.GetClaimsFromToken(helper.GetBearerTokenValue(c))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var articleRequestDTO model.ArticleRequestDTO
		if err := c.ShouldBindJSON(&articleRequestDTO); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if claims.Role == "REGULAR" && (articleRequestDTO.Status != string(model.DRAFT) &&
			articleRequestDTO.Status != string(model.PENDING)) {
			articleRequestDTO.Status = string(model.DRAFT)
		}

		createdArticle, err := ac.articleService.CreateArticle(articleRequestDTO, claims.UserId)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		c.JSON(http.StatusCreated, createdArticle.GetArticleDTO())
	}
}

func (ac *articleController) GetArticle() gin.HandlerFunc {
	return func(c *gin.Context) {
		articleId := uuid.MustParse(c.Param("article_id"))
		foundArticle, err := ac.articleService.GetArticleById(articleId)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		c.JSON(http.StatusCreated, foundArticle.GetArticleDTO())
	}
}

func (ac *articleController) GetArticlesByStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		articleStatus := c.Param("status")
		recordPerPage, err := strconv.Atoi(c.DefaultQuery("recordPerPage", "20"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid page size param"})
			return
		}
		page, err := strconv.Atoi(c.DefaultQuery("pageNumber", "1"))
		if err != nil || page < 1 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid page number param"})
			return
		}

		articles, totalCount, err := ac.articleService.GetArticlesByStatus(
			recordPerPage,
			page,
			model.ArticleStatus(articleStatus),
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while retrieving articles"})
			return
		}

		shortArticlesDTO := ConvertArticlesToShortArticlesDTO(*articles)

		c.JSON(http.StatusOK, gin.H{
			"articles":   shortArticlesDTO,
			"totalCount": totalCount,
		})
	}
}

func (ac *articleController) GetMyArticlesByStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		articleStatus := c.Param("status")
		claims, err := helper.GetClaimsFromToken(helper.GetBearerTokenValue(c))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		recordPerPage, err := strconv.Atoi(c.DefaultQuery("recordPerPage", "20"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid page size param"})
			return
		}
		page, err := strconv.Atoi(c.DefaultQuery("pageNumber", "1"))
		if err != nil || page < 1 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid page number param"})
			return
		}

		articles, totalCount, err := ac.articleService.GetArticlesByStatusByUserId(
			recordPerPage,
			page,
			model.ArticleStatus(articleStatus),
			uuid.MustParse(claims.UserId),
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while retrieving articles"})
			return
		}

		shortArticlesDTO := ConvertArticlesToShortArticlesDTO(*articles)

		c.JSON(http.StatusOK, gin.H{
			"articles":   shortArticlesDTO,
			"totalCount": totalCount,
		})
	}
}

func (ac *articleController) SetArticleStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		var statusQuery struct {
			ArticleId string `json:"article_id" validate:"required"`
			Status    string `json:"status" validate:"required"`
		}
		if err := c.ShouldBindJSON(&statusQuery); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request data"})
			return
		}
		updatedArticle, err := ac.articleService.SetArticleStatus(
			uuid.MustParse(statusQuery.ArticleId),
			statusQuery.Status,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, updatedArticle.GetShortArticleDTO())
	}
}

func (ac *articleController) GetArticleComments() gin.HandlerFunc {
	return func(c *gin.Context) {
		articleId := uuid.MustParse(c.Param("article_id"))
		recordPerPage, err := strconv.Atoi(c.DefaultQuery("recordPerPage", "20"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid page size param"})
			return
		}
		page, err := strconv.Atoi(c.DefaultQuery("pageNumber", "1"))
		if err != nil || page < 1 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid page number param"})
			return
		}

		comments, totalCount, err := ac.commentService.GetArticleComments(
			recordPerPage,
			page,
			articleId,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while retrieving article comments"})
			return
		}

		shortCommentsDTO := ConvertCommentsToShortCommentsDTO(*comments)

		c.JSON(http.StatusOK, gin.H{
			"articles":   shortCommentsDTO,
			"totalCount": totalCount,
		})

	}
}

func (ac *articleController) CreateComment() gin.HandlerFunc {
	return func(c *gin.Context) {
		var comReqDTO model.CommentRequestDTO
		err := c.ShouldBindJSON(&comReqDTO)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		claims, err := helper.GetClaimsFromToken(helper.GetBearerTokenValue(c))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		comment, err := ac.commentService.CreateComment(comReqDTO, claims.UserId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, comment.GetShortCommentDTO())
	}
}

func (ac *articleController) SetCommentIsVisible() gin.HandlerFunc {
	return func(c *gin.Context) {
		var comVisibilityReq struct {
			CommentId string `json:"comment_id"`
			IsVisible bool   `json:"is_visible"`
		}
		if err := c.ShouldBindJSON(&comVisibilityReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		comment, err := ac.commentService.SetCommentIsVisible(
			comVisibilityReq.CommentId,
			comVisibilityReq.IsVisible,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, comment.GetCommentDTO())
	}
}

func ConvertArticlesToShortArticlesDTO(articles []model.Article) []model.ShortArticleDTO {
	var converted []model.ShortArticleDTO = make([]model.ShortArticleDTO, len(articles))
	for i := range articles {
		converted[i] = *articles[i].GetShortArticleDTO()
	}
	return converted
}

func ConvertCommentsToShortCommentsDTO(comments []model.Comment) []model.ShortCommentDTO {
	var converted []model.ShortCommentDTO = make([]model.ShortCommentDTO, len(comments))
	for i := range comments {
		converted[i] = *comments[i].GetShortCommentDTO()
	}
	return converted
}
