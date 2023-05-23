package repository

import (
	"heptaber/auth/domain/model"

	"github.com/google/uuid"
)

type IUserRepository interface {
	Save(user *model.User) error
	FindByID(id uuid.UUID) (*model.User, error)
	SetUserVerifiedTrueByID(id uuid.UUID) error
	SetRegularUserIsLockedByUserId(id uuid.UUID, isLocked bool) (*model.User, error)
	SetUserRoleById(userId uuid.UUID, role string) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	GetAllUsersByPage(size int, num int) (*[]model.User, int64, error)
	Update(user *model.User) error
	Delete(id uuid.UUID) error
}
