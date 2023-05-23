package repository

import (
	"heptaber/blog/app/helper"
	"heptaber/blog/domain/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *commentRepository {
	return &commentRepository{db: db}
}

func (cr *commentRepository) Save(comment *model.Comment) error {
	return cr.db.Create(comment).Error
}

func (cr *commentRepository) FindAllByArticleId(articleId uuid.UUID) (*[]model.Comment, error) {
	var comments *[]model.Comment
	if err := cr.db.Model(&model.Comment{}).
		Where("article_id = ? AND status = ?", articleId, model.POSTED).
		Find(comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

func (cr *commentRepository) FindAllCommentsByArticleIdByPage(pageSize int, pageNum int, articleId uuid.UUID) (
	*[]model.Comment, int64, error,
) {

	var comments *[]model.Comment
	var totalCount int64 = 0
	var err error = cr.db.Model(&model.Comment{}).
		Where("article_id = ? AND status = ?", articleId, model.POSTED).
		Count(&totalCount).
		Offset((pageNum - 1) * pageSize).
		Limit(pageSize).
		Order("created_at DESC").
		Find(&comments).
		Error
	return comments, totalCount, err
}

func (cr *commentRepository) Update(comment *model.Comment) error {
	return cr.db.Save(comment).Error
}

func (cr *commentRepository) SetCommentIsVisibleById(commentId uuid.UUID, isVisible bool) (*model.Comment, error) {
	var comment *model.Comment
	if err := cr.db.Model(&model.Comment{}).
		Where("id = ?", commentId).
		Updates(map[string]interface{}{
			"is_visible": isVisible,
			"updated_at": helper.GetUTCCurrentTimeRFC3339(),
		}).
		First(&comment).Error; err != nil {
		return nil, err
	}
	return comment, nil
}
