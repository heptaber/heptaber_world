package repository

import (
	"heptaber/auth/domain/model"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type verificationCodeRepository struct {
	db *gorm.DB
}

func NewVerificationCodeRepository(db *gorm.DB) *verificationCodeRepository {
	return &verificationCodeRepository{db: db}
}

func (vcr *verificationCodeRepository) Save(verificationCode *model.VerificationCode) error {
	return vcr.db.Create(verificationCode).Error
}

func (vcr *verificationCodeRepository) FindById(id uuid.UUID) (*model.VerificationCode, error) {
	vCode := &model.VerificationCode{}
	if err := vcr.db.Model(&model.VerificationCode{}).
		Where("id = ?", id).
		First(vCode).Error; err != nil {
		return nil, err
	}
	return vCode, nil
}

func (vcr *verificationCodeRepository) FindByUserId(userId uuid.UUID) (*model.VerificationCode, error) {
	vCode := &model.VerificationCode{}
	if err := vcr.db.Model(&model.VerificationCode{}).
		Where("user_id = ?", userId).
		First(vCode).Error; err != nil {
		return nil, err
	}
	return vCode, nil
}

func (vcr *verificationCodeRepository) DeleteAllExpired() error {
	currentTime := time.Now().UTC()
	return vcr.db.Where("expires_at < ?", currentTime).
		Delete(&model.VerificationCode{}).Error
}

func (vcr *verificationCodeRepository) DeleteById(id uuid.UUID) error {
	return vcr.db.Delete(&model.VerificationCode{}, id).Error
}
