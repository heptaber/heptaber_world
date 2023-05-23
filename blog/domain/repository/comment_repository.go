package repository

import (
	"heptaber/blog/domain/model"

	"github.com/google/uuid"
)

type ICommentRepository interface {
	Save(comment *model.Comment) error
	FindAllByArticleId(articleId uuid.UUID) (*[]model.Comment, error)
	FindAllCommentsByArticleIdByPage(pageSize int, pageNum int, articleId uuid.UUID) (
		*[]model.Comment, int64, error,
	)
	Update(article *model.Comment) error
	SetCommentIsVisibleById(commentId uuid.UUID, isVisible bool) (*model.Comment, error)
}
