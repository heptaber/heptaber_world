package model

import (
	"time"

	"github.com/google/uuid"
)

type VerificationCode struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	CreatedAt time.Time `gorm:"not null"`
	ExpiresAt time.Time `gorm:"not null"`
	UserID    uuid.UUID `gorm:"type:uuid;not null"`
}
