package repository

import (
	"heptaber/blog/domain/model"

	"github.com/google/uuid"
)

type IArticleRepository interface {
	Save(article *model.Article) error
	FindByID(id uuid.UUID) (*model.Article, error)
	SetArticleStatus(id uuid.UUID, status string) (*model.Article, error)
	FindAllArticlesByStatusByPage(pageSize int, pageNum int, status string,
	) (*[]model.Article, int64, error)
	FindAllArticlesByAuthorIDByStatusByPage(pageSize int, pageNum int, status string, authorId uuid.UUID,
	) (*[]model.Article, int64, error)
	Update(article *model.Article) error
}
