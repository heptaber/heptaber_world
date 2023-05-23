package service

import (
	"heptaber/blog/app/helper"
	"heptaber/blog/domain/model"
	"heptaber/blog/domain/repository"
	"log"

	"github.com/google/uuid"
)

type articleService struct {
	articleRepo repository.IArticleRepository
}

func NewArticleService(articleRepo repository.IArticleRepository) *articleService {
	return &articleService{articleRepo: articleRepo}
}

func (as *articleService) CreateArticle(articleRequestDTO model.ArticleRequestDTO, authorId string) (*model.Article, error) {
	author_id := uuid.MustParse(authorId)
	currentTime := helper.GetUTCCurrentTimeRFC3339()
	article := &model.Article{
		Title:       articleRequestDTO.Title,
		Description: articleRequestDTO.Description,
		Content:     articleRequestDTO.Content,
		AuthorID:    author_id,
		Status:      model.ArticleStatus(articleRequestDTO.Status),
		CreatedAt:   currentTime,
		UpdatedAt:   currentTime,
	}
	if articleRequestDTO.Status == string(model.POSTED) {
		article.PostedAt = currentTime
	}
	err := as.articleRepo.Save(article)
	if err != nil {
		log.Fatal("error while saving article in database: ", err.Error())
		return nil, err
	}
	return article, nil
}

func (as *articleService) GetArticleById(id uuid.UUID) (*model.Article, error) {
	return as.articleRepo.FindByID(id)
}

func (as *articleService) GetArticlesByStatus(
	recordPerPage int, page int, status model.ArticleStatus,
) (*[]model.Article, int64, error) {
	return as.articleRepo.FindAllArticlesByStatusByPage(recordPerPage, page, string(status))
}

func (as *articleService) GetArticlesByStatusByUserId(
	recordPerPage int, page int, status model.ArticleStatus, authorId uuid.UUID,
) (*[]model.Article, int64, error) {
	return as.articleRepo.FindAllArticlesByAuthorIDByStatusByPage(recordPerPage, page, string(status), authorId)
}

func (as *articleService) SetArticleStatus(articleId uuid.UUID, status string) (*model.Article, error) {
	return as.articleRepo.SetArticleStatus(articleId, status)
}
