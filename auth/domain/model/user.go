package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID               uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Email            string    `gorm:"unique;not null"`
	Password         string    `gorm:"not null"`
	Username         string    `gorm:"type:varchar(255);not null"`
	Role             UserRole  `gorm:"type:user_role"`
	CreatedAt        time.Time `gorm:"not null"`
	UpdatedAt        time.Time `gorm:"not null"`
	DeletedAt        time.Time
	Token            *Token            `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	VerificationCode *VerificationCode `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	IsVerified       bool              `gorm:"not null;default:false"`
	IsLocked         bool              `gorm:"not null;default:false"`
}

type UserRole string

const (
	ADMIN     UserRole = "ADMIN"
	MODERATOR UserRole = "MODERATOR"
	REGULAR   UserRole = "REGULAR"
)

type UserDTO struct {
	ID         string    `json:"id,omitempty"`
	Email      string    `json:"email" validate:"required,min=5,max=150"`
	Username   string    `json:"username" validate:"required,min=2,max=100"`
	Role       string    `json:"role" validate:"required,eq=ADMIN|eq=MODERATOR|eq=REGULAR"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
	DeletedAt  time.Time `json:"deleted_at,omitempty"`
	IsVerified *bool     `json:"is_verified,omitempty"`
	IsLocked   *bool     `json:"is_locked,omitempty"`
}

type ShortUserDTO struct {
	ID       string `json:"id,omitempty"`
	Email    string `json:"email" validate:"required,min=5,max=150"`
	Username string `json:"username" validate:"required,min=2,max=100"`
	Role     string `json:"role" validate:"required,eq=ADMIN|eq=MODERATOR|eq=REGULAR"`
}

func (u *User) GetUserDTO() *UserDTO {
	return &UserDTO{
		ID:         u.ID.String(),
		Email:      u.Email,
		Username:   u.Username,
		Role:       string(u.Role),
		CreatedAt:  u.CreatedAt,
		UpdatedAt:  u.UpdatedAt,
		DeletedAt:  u.DeletedAt,
		IsVerified: &u.IsVerified,
		IsLocked:   &u.IsLocked,
	}
}

func (u *User) GetShortUserDTO() *ShortUserDTO {
	return &ShortUserDTO{
		ID:       u.ID.String(),
		Email:    u.Email,
		Username: u.Username,
		Role:     string(u.Role),
	}
}
