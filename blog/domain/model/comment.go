package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Comment struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Content   string         `gorm:"not null"`
	AuthorID  uuid.UUID      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	IsVisible bool           `gorm:"not null;default:true"`
	ArticleID uuid.UUID      `gorm:"index;column:article_id"`
	CreatedAt time.Time      `gorm:"not null"`
	UpdatedAt time.Time      `gorm:"not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type CommentDTO struct {
	ID        string         `json:"id"`
	Content   string         `json:"content"`
	AuthorID  string         `json:"author_id"`
	IsVisible *bool          `json:"is_visible"`
	ArticleID string         `json:"article_id"`
	CreatedAt time.Time      `json:"created_at,omitempty"`
	UpdatedAt time.Time      `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type ShortCommentDTO struct {
	ID        string    `json:"id"`
	Content   string    `json:"content"`
	AuthorID  string    `json:"author_id"`
	ArticleID string    `json:"article_id"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type CommentRequestDTO struct {
	Content   string `json:"content" validate:"required,min=1,max=1024"`
	ArticleId string `json:"article_id" validate:"required"`
}

func (c *Comment) GetShortCommentDTO() *ShortCommentDTO {
	return &ShortCommentDTO{
		ID:        c.ID.String(),
		Content:   c.Content,
		AuthorID:  c.AuthorID.String(),
		ArticleID: c.ArticleID.String(),
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}

func (c *Comment) GetCommentDTO() *CommentDTO {
	return &CommentDTO{
		ID:        c.ID.String(),
		Content:   c.Content,
		AuthorID:  c.AuthorID.String(),
		IsVisible: &c.IsVisible,
		ArticleID: c.ArticleID.String(),
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
		DeletedAt: c.DeletedAt,
	}
}
