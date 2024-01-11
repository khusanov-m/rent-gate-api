package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Subscription struct {
	ID                 uint      `gorm:"primaryKey;autoIncrement"`
	UUID               uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();uniqueIndex"`
	UserID             uint      `gorm:"not null"`
	SubscriptionTypeID string    `gorm:"not null"`
	StartDate          time.Time `gorm:"not null"`
	EndDate            time.Time `gorm:"not null"`
	TotalPrice         float64   `gorm:"not null"`

	CreatedAt time.Time      `gorm:"not null"`
	UpdatedAt time.Time      `gorm:"not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	//User *User `gorm:"foreignkey:UserID" json:"user,omitempty"`
}
