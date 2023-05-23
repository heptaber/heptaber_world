package repository

import (
	"fmt"
	"heptaber/auth/app/helper"
	"heptaber/auth/domain/model"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Save(user *model.User) error {
	err := r.db.Create(user).Error
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return fmt.Errorf("such user already exists")
		}
		return err
	}
	return nil
}

func (r *userRepository) FindByID(id uuid.UUID) (*model.User, error) {
	user := &model.User{}
	if err := r.db.Model(&model.User{}).
		Where("id = ?", id).
		First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) SetUserVerifiedTrueByID(id uuid.UUID) error {
	if err := r.db.Model(&model.User{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"is_verified": true,
			"updated_at":  helper.GetUTCCurrentTimeRFC3339(),
		}).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) SetRegularUserIsLockedByUserId(id uuid.UUID, isLocked bool) (*model.User, error) {
	user := &model.User{}
	if err := r.db.Model(&model.User{}).
		Where("id = ? AND role = 'REGULAR'", id).
		Updates(map[string]interface{}{
			"is_locked":  isLocked,
			"updated_at": helper.GetUTCCurrentTimeRFC3339(),
		}).
		First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) SetUserRoleById(id uuid.UUID, role string) (*model.User, error) {
	user := &model.User{}
	if err := r.db.Model(&model.User{}).
		Where("id = ? AND role != 'ADMIN'", id).
		UpdateColumns(map[string]interface{}{
			"role":       role,
			"updated_at": helper.GetUTCCurrentTimeRFC3339(),
		}).
		First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) FindByEmail(email string) (*model.User, error) {
	user := &model.User{}
	if err := r.db.Model(&model.User{}).
		Where("email = ?", email).
		First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) GetAllUsersByPage(pageSize int, pageNum int) (users *[]model.User, totalCount int64, err error) {
	totalCount = 0
	err = r.db.Model(&model.User{}).
		Count(&totalCount).
		Offset((pageNum - 1) * pageSize).
		Limit(pageSize).
		Order("username ASC").
		Find(&users).
		Error
	return
}

func (r *userRepository) Update(user *model.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) Delete(id uuid.UUID) error {
	// TODO: implement more safe version :)
	return r.db.Delete(&model.User{}, id).Error
}
