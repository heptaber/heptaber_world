package repository

import (
	"heptaber/auth/domain/model"

	"github.com/google/uuid"
)

type ITokenRepository interface {
	Save(token *model.Token) error
	UpdateToken(token *model.Token) error
	FindByID(id uuid.UUID) (*model.Token, error)
	FindByUserID(id uuid.UUID) (*model.Token, error)
	DeleteByID(id uuid.UUID) error
	DeleteByUserID(id uuid.UUID) error
}
