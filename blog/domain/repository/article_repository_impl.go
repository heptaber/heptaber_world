package repository

import (
	"heptaber/blog/app/helper"
	"heptaber/blog/domain/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) *articleRepository {
	return &articleRepository{db: db}
}

func (r *articleRepository) Save(article *model.Article) error {
	return r.db.Create(article).Error
}

func (r *articleRepository) FindByID(id uuid.UUID) (*model.Article, error) {
	article := &model.Article{}
	if err := r.db.Model(&model.Article{}).
		Where("id = ?", id).
		First(article).Error; err != nil {
		return nil, err
	}
	return article, nil
}

func (r *articleRepository) SetArticleStatus(id uuid.UUID, status string) (*model.Article, error) {
	article := &model.Article{}
	if err := r.db.Model(&model.Article{}).
		Where("id = ?", id).
		UpdateColumns(map[string]interface{}{
			"status":     status,
			"updated_at": helper.GetUTCCurrentTimeRFC3339(),
		}).First(article).Error; err != nil {
		return nil, err
	}
	return article, nil
}

func (r *articleRepository) FindAllArticlesByStatusByPage(pageSize int, pageNum int, status string,
) (*[]model.Article, int64, error) {

	var articles *[]model.Article
	var totalCount int64 = 0
	var err error = r.db.Model(&model.Article{}).
		Where("status = ?", status).
		Count(&totalCount).
		Offset((pageNum - 1) * pageSize).
		Limit(pageSize).
		Order("posted_at DESC").
		Find(&articles).
		Error
	return articles, totalCount, err
}

func (r *articleRepository) FindAllArticlesByAuthorIDByStatusByPage(
	pageSize int,
	pageNum int,
	status string,
	authorId uuid.UUID,
) (*[]model.Article, int64, error) {

	var articles *[]model.Article
	var totalCount int64 = 0
	var err error = r.db.Model(&model.Article{}).
		Where("author_id = ? AND status = ?", authorId, status).
		Count(&totalCount).
		Offset((pageNum - 1) * pageSize).
		Limit(pageSize).
		Order("posted_at DESC").
		Find(&articles).
		Error
	return articles, totalCount, err
}

func (r *articleRepository) Update(article *model.Article) error {
	return r.db.Save(article).Error
}
