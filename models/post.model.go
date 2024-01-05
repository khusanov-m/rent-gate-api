package models

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	UUID      uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();uniqueIndex" json:"uuid,omitempty"`
	Title     string    `gorm:"uniqueIndex;not null" json:"title,omitempty"`
	Content   string    `gorm:"not null" json:"content,omitempty"`
	Image     string    `gorm:"not null" json:"image,omitempty"`
	CreatedAt time.Time `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at,omitempty"`
	UserID    uint      `gorm:"not null" json:"user_id,omitempty"`
	// User      *User     `gorm:"foreignkey:UserID" json:"user,omitempty"`
}

type CreatePostInput struct {
	Title     string    `json:"title"  binding:"required"`
	Content   string    `json:"content" binding:"required"`
	Image     string    `json:"image" binding:"required"`
	User      string    `json:"user,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type UpdatePostInput struct {
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
	Image   string `json:"image,omitempty"`
}

type PostResponse struct {
	UUID    uuid.UUID `json:"uuid,omitempty"`
	Title   string    `json:"title,omitempty"`
	Content string    `json:"content,omitempty"`
	Image   string    `json:"image,omitempty"`
	User    *User     `json:"user,omitempty"`
}
