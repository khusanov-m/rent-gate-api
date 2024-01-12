package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID                 uint      `gorm:"primaryKey;autoIncrement"`
	UUID               uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();uniqueIndex"`
	Name               string    `gorm:"type:varchar(255);default:uuid_generate_v4();not null"`
	Email              string    `gorm:"type:varchar(255);uniqueIndex;not null"`
	Password           string    `gorm:"not null"`
	Role               string    `gorm:"type:varchar(255);default:'user';not null"`
	Provider           string    `gorm:"type:varchar(255);not null"`
	Photo              string    `gorm:"type:varchar(255);not null"`
	VerificationCode   string    `gorm:"type:varchar(255)"`
	PasswordResetToken string    `gorm:"type:varchar(255)"`
	PasswordResetAt    time.Time
	Verified           bool `gorm:"default:false;not null"`

	Subscription   *Subscription   `gorm:"foreignkey:UserID"`
	LoyaltyAccount *LoyaltyAccount `gorm:"foreignkey:UserID"`

	CreatedAt time.Time      `gorm:"not null"`
	UpdatedAt time.Time      `gorm:"not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
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
	ID       uuid.UUID `json:"id,omitempty"`
	Name     string    `json:"name,omitempty"`
	Email    string    `json:"email,omitempty"`
	Role     string    `json:"role,omitempty"`
	Photo    string    `json:"photo,omitempty"`
	Provider string    `json:"provider"`
	Verified bool      `json:"verified"`

	Subscription   *Subscription   `json:"subscription,omitempty"`
	LoyaltyAccount *LoyaltyAccount `json:"loyalty_account,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type ForgotPasswordInput struct {
	Email string `json:"email" binding:"required"`
}

type ResetPasswordInput struct {
	Password        string `json:"password" binding:"required"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required"`
}
