package repository

import (
	"heptaber/auth/domain/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type tokenRepository struct {
	db *gorm.DB
}

func NewTokenRepository(db *gorm.DB) *tokenRepository {
	return &tokenRepository{db: db}
}

func (r *tokenRepository) Save(token *model.Token) error {
	return r.db.Create(token).Error
}

func (r *tokenRepository) UpdateToken(token *model.Token) error {
	if err := r.db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "user_id"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"refresh_token",
			"expires_at",
			"updated_at",
		}),
	}).Create(&token).Error; err != nil {
		return err
	}
	return nil
}

func (r *tokenRepository) FindByID(id uuid.UUID) (*model.Token, error) {
	token := &model.Token{}
	if err := r.db.Model(&model.Token{}).
		Where("id = ?", id).
		First(token).Error; err != nil {
		return nil, err
	}
	return token, nil
}

func (r *tokenRepository) FindByUserID(userId uuid.UUID) (*model.Token, error) {
	token := &model.Token{}
	if err := r.db.Model(&model.Token{}).
		Where("user_id = ?", userId).
		First(token).Error; err != nil {
		return nil, err
	}
	return token, nil
}

func (r *tokenRepository) DeleteByID(id uuid.UUID) error {
	return r.db.Delete(&model.Token{}, id).Error
}

func (r *tokenRepository) DeleteByUserID(id uuid.UUID) error {
	return r.db.Delete(&model.Token{}, "user_id = ?", id).Error
}
