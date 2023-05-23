package model

import (
	"time"

	"github.com/google/uuid"
)

type NotificationType string

const (
	EMAIL        NotificationType = "email"
	NOTIFICATION NotificationType = "notification"
)

type EventType string

const (
	VERIFICATION EventType = "verification"
	SUBSCRIPTION EventType = "subscription"
)

type VerificationEmail struct {
	Recipient        string `json:"recipient"`
	VerificationCode string `json:"code"`
}

// TODO: impl
type Notification struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
	IsViewed  bool      `gorm:"is_viewed;not null;default:false"`
	UserId    uuid.UUID `gorm:"type:uuid;not null"`
	DeletedAt time.Time
}

type Message struct {
	EventType string `json:"notification_type"`
	Payload   string `json:"payload"`
}
