package controllers

import "github.com/gin-gonic/gin"

type IArticleController interface {
	GetArticles() gin.HandlerFunc
	CreateArticle() gin.HandlerFunc
	GetArticle() gin.HandlerFunc
	GetArticlesByStatus() gin.HandlerFunc
	GetMyArticlesByStatus() gin.HandlerFunc
	SetArticleStatus() gin.HandlerFunc
	GetArticleComments() gin.HandlerFunc
	CreateComment() gin.HandlerFunc
	SetCommentIsVisible() gin.HandlerFunc
	// TODO: add find by title articles
}
