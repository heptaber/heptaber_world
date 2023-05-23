package service

import (
	"heptaber/auth/domain/model"
)

type IUserService interface {
	GetUsers(pageSize int, pageNum int) (*[]model.User, int64, error)
	GetUserById(userId string) (*model.User, error)
	GetUserByEmail(userEmail string) (*model.User, error)
	SetUserRoleByUserId(userId string, role string) (*model.User, error)
	SetUserLockByUserId(userId string, isLocked bool) (*model.User, error)
}
