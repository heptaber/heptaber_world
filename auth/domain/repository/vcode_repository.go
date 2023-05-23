package repository

import (
	"heptaber/auth/domain/model"

	"github.com/google/uuid"
)

type IVerificationCodeRepository interface {
	Save(*model.VerificationCode) error
	FindById(id uuid.UUID) (*model.VerificationCode, error)
	FindByUserId(userId uuid.UUID) (*model.VerificationCode, error)
	DeleteAllExpired() error
	DeleteById(id uuid.UUID) error
}
