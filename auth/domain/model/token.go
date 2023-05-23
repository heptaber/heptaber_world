package model

import (
	"time"

	"github.com/google/uuid"
)

type Token struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	UserID       uuid.UUID `gorm:"type:uuid;not null;unique"`
	RefreshToken string    `gorm:"not null"`
	ExpiresAt    time.Time `gorm:"not null"`
	CreatedAt    time.Time `gorm:"not null"`
	UpdatedAt    time.Time `gorm:"not null"`
}
