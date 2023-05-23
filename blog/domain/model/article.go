package model

import (
	"time"

	"github.com/google/uuid"
)

type Article struct {
	ID          uuid.UUID     `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Title       string        `gorm:"not null"`
	Description string        `gorm:"default:''"`
	Content     string        `gorm:"not null"`
	AuthorID    uuid.UUID     `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Status      ArticleStatus `gorm:"type:article_status;not null"`
	Comments    []Comment     `gorm:"foreignKey:ArticleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt   time.Time     `gorm:"not null"`
	UpdatedAt   time.Time     `gorm:"not null"`
	PostedAt    time.Time
	DeletedAt   time.Time
}

type ArticleStatus string

const (
	PENDING  ArticleStatus = "PENDING"
	POSTED   ArticleStatus = "POSTED"
	REJECTED ArticleStatus = "REJECTED"
	DRAFT    ArticleStatus = "DRAFT"
)

type ArticleRequestDTO struct {
	Title       string `json:"title" validate:"required,min=3,max=512"`
	Description string `json:"description"`
	Content     string `json:"content" validate:"required,min=100,max=4096"`
	Status      string `json:"status" validate:"required,eq=PENDING|eq=POSTED|eq=REJECTED|eq=DRAFT"`
}

type ArticleDTO struct {
	ID          string    `json:"id,omitempty"`
	Title       string    `json:"title" validate:"required,min=3,max=512"`
	Description string    `json:"description"`
	Content     string    `json:"content" validate:"required,min=100,max=4096"`
	AuthorID    string    `json:"author_id,omitempty"`
	Status      string    `json:"status" validate:"required,eq=PENDING|eq=POSTED|eq=REJECTED|eq=DRAFT"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	DeletedAt   time.Time `json:"deleted_at,omitempty"`
	PostedAt    time.Time `json:"posted_at,omitempty"`
}

type ShortArticleDTO struct {
	ID          string    `json:"id,omitempty"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	AuthorID    string    `json:"author_id"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	PostedAt    time.Time `json:"posted_at,omitempty"`
}

func (a *Article) GetArticleDTO() *ArticleDTO {
	return &ArticleDTO{
		ID:          a.ID.String(),
		Title:       a.Title,
		Description: a.Description,
		Content:     a.Content,
		AuthorID:    a.AuthorID.String(),
		Status:      string(a.Status),
		CreatedAt:   a.CreatedAt,
		UpdatedAt:   a.UpdatedAt,
		DeletedAt:   a.DeletedAt,
		PostedAt:    a.PostedAt,
	}
}

func (a *Article) GetShortArticleDTO() *ShortArticleDTO {
	return &ShortArticleDTO{
		ID:          a.ID.String(),
		Title:       a.Title,
		Description: a.Description,
		AuthorID:    a.AuthorID.String(),
		Status:      string(a.Status),
		CreatedAt:   a.CreatedAt,
		UpdatedAt:   a.UpdatedAt,
		PostedAt:    a.PostedAt,
	}
}
