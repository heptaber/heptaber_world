package service

import (
	"heptaber/blog/domain/model"

	"github.com/google/uuid"
)

type ICommentService interface {
	CreateComment(commentRequestDTO model.CommentRequestDTO, authorId string) (*model.Comment, error)
	GetArticleComments(
		recordPerPage int,
		page int,
		articleId uuid.UUID,
	) (*[]model.Comment, int64, error)
	SetCommentIsVisible(commentId string, isVisible bool) (*model.Comment, error)
}
