package service

import (
	"heptaber/blog/domain/model"

	"github.com/google/uuid"
)

type IArticleService interface {
	CreateArticle(articleRequestDTO model.ArticleRequestDTO, authorId string) (*model.Article, error)
	GetArticleById(id uuid.UUID) (*model.Article, error)
	GetArticlesByStatus(recordPerPage int, page int, status model.ArticleStatus,
	) (*[]model.Article, int64, error)
	GetArticlesByStatusByUserId(
		recordPerPage int,
		page int,
		status model.ArticleStatus,
		authorId uuid.UUID,
	) (*[]model.Article, int64, error)
	SetArticleStatus(articleId uuid.UUID, status string) (*model.Article, error)
}
