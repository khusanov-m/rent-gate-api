package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// PostSchema
type Post struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"id,omitempty"`
	UUID      uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();uniqueIndex" json:"uuid,omitempty"`
	Title     string         `gorm:"uniqueIndex;not null"`
	Content   string         `gorm:"not null"`
	Image     string         `gorm:"not null"`
	CreatedAt time.Time      `gorm:"not null"`
	UpdatedAt time.Time      `gorm:"not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	UserID uint
	//User   *User `gorm:"foreignkey:UserID" json:"user,omitempty"`
}

type CreatePostInput struct {
	Title   string `json:"title"  binding:"required"`
	Content string `json:"content" binding:"required"`
	Image   string `json:"image" binding:"required"`
}

type UpdatePostInput struct {
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
	Image   string `json:"image,omitempty"`
}

type PostResponse struct {
	ID      uuid.UUID     `json:"id,omitempty"`
	Title   string        `json:"title,omitempty"`
	Content string        `json:"content,omitempty"`
	Image   string        `json:"image,omitempty"`
	User    *UserResponse `json:"user,omitempty"`
}
