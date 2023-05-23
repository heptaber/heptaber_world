package service

import (
	"errors"
	"heptaber/auth/domain/model"
	"heptaber/auth/domain/repository"

	"github.com/google/uuid"
)

type userService struct {
	repository.IUserRepository
}

func NewUserService(repository repository.IUserRepository) *userService {
	return &userService{IUserRepository: repository}
}

func (us *userService) GetUsers(pageSize int, pageNum int) (users *[]model.User, totalCount int64, err error) {
	return us.IUserRepository.GetAllUsersByPage(pageSize, pageNum)
}

func (us *userService) GetUserById(userId string) (*model.User, error) {
	user_id := uuid.MustParse(userId)
	return us.IUserRepository.FindByID(user_id)
}

func (us *userService) GetUserByEmail(userEmail string) (foundUser *model.User, err error) {
	return us.IUserRepository.FindByEmail(userEmail)
}

func (us *userService) SetUserRoleByUserId(userId string, role string) (updatedUser *model.User, err error) {
	if string(model.MODERATOR) != role && string(model.REGULAR) != role {
		return nil, errors.New("there is no such role")
	}
	user_id := uuid.MustParse(userId)
	updatedUser, err = us.IUserRepository.SetUserRoleById(user_id, role)
	if err != nil {
		return nil, err
	}

	return
}

func (us *userService) SetUserLockByUserId(userId string, isLocked bool) (*model.User, error) {
	user_id := uuid.MustParse(userId)
	user, err := us.IUserRepository.SetRegularUserIsLockedByUserId(user_id, isLocked)
	if err != nil {
		return nil, err
	}

	return user, nil
}
