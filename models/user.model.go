package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID                 uint      `gorm:"primaryKey;autoIncrement"`
	UUID               uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();uniqueIndex"`
	Name               string    `gorm:"type:varchar(255);not null"`
	Email              string    `gorm:"type:varchar(255);uniqueIndex;not null"`
	Password           string    `gorm:"not null"`
	Role               string    `gorm:"type:varchar(255);not null"`
	Provider           string    `gorm:"type:varchar(255);not null"`
	Photo              string    `gorm:"type:varchar(255);not null"`
	VerificationCode   string    `gorm:"type:varchar(255)"`
	PasswordResetToken string    `gorm:"type:varchar(255)"`
	PasswordResetAt    time.Time
	Verified           bool           `gorm:"not null"`
	CreatedAt          time.Time      `gorm:"not null"`
	UpdatedAt          time.Time      `gorm:"not null"`
	DeletedAt          gorm.DeletedAt `gorm:"index"`
	Vehicles           []Vehicle      `gorm:"foreignkey:UserID"`
	// Subscription       *Subscription   `gorm:"foreignkey:UserID"`
	LoyaltyProgram     *LoyaltyProgram `gorm:"foreignkey:UserID"`
	Posts []Post `gorm:"foreignkey:UserID"`
}

type SignUpInput struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required,min=8"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required"`
	Photo           string `json:"photo" binding:"required"`
}

type SignInInput struct {
	Email    string `json:"email"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}

type UserResponse struct {
	ID             uuid.UUID         `json:"id,omitempty"`
	Name           string            `json:"name,omitempty"`
	Email          string            `json:"email,omitempty"`
	Role           string            `json:"role,omitempty"`
	Photo          string            `json:"photo,omitempty"`
	Provider       string            `json:"provider"`
	CreatedAt      time.Time         `json:"created_at,omitempty"`
	UpdatedAt      time.Time         `json:"updated_at,omitempty"`
	Verified       bool              `json:"verified"`
	Vehicles       []VehicleResponse `json:"vehicles,omitempty"`
	// Subscription   *Subscription     `json:"subscription,omitempty"`
	LoyaltyProgram *LoyaltyProgram   `json:"loyalty_program,omitempty"`
	Posts          []PostResponse    `json:"posts,omitempty"`
}

// ? ForgotPasswordInput struct
type ForgotPasswordInput struct {
	Email string `json:"email" binding:"required"`
}

// ? ResetPasswordInput struct
type ResetPasswordInput struct {
	Password        string `json:"password" binding:"required"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required"`
}
