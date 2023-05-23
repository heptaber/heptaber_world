package service

import (
	"heptaber/blog/app/helper"
	"heptaber/blog/domain/model"
	"heptaber/blog/domain/repository"
	"log"

	"github.com/google/uuid"
)

type commentService struct {
	commentRepo repository.ICommentRepository
}

func NewCommentService(commentRepository repository.ICommentRepository) *commentService {
	return &commentService{commentRepo: commentRepository}
}

func (cs *commentService) CreateComment(commentRequestDTO model.CommentRequestDTO, authorId string) (*model.Comment, error) {
	author_id := uuid.MustParse(authorId)
	article_id := uuid.MustParse(commentRequestDTO.ArticleId)
	currentTime := helper.GetUTCCurrentTimeRFC3339()
	comment := &model.Comment{
		Content:   commentRequestDTO.Content,
		ArticleID: article_id,
		AuthorID:  author_id,
		IsVisible: true,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}
	err := cs.commentRepo.Save(comment)
	if err != nil {
		log.Fatal("error while saving comment in database: ", err.Error())
		return nil, err
	}
	return comment, nil
}

func (cs *commentService) GetArticleComments(recordPerPage int, page int, articleId uuid.UUID,
) (*[]model.Comment, int64, error) {
	return cs.commentRepo.FindAllCommentsByArticleIdByPage(recordPerPage, page, articleId)
}

func (cs *commentService) SetCommentIsVisible(commentId string, isVisible bool,
) (*model.Comment, error) {
	comment_id := uuid.MustParse(commentId)
	return cs.commentRepo.SetCommentIsVisibleById(comment_id, isVisible)
}
