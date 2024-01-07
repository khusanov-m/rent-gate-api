package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Subscription struct {
	ID         uint      `gorm:"primaryKey;autoIncrement"`
	UUID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();uniqueIndex"`
	UserID     uint
	StartDate  time.Time
	EndDate    time.Time
	MonthlyFee float64
	CreatedAt  time.Time      `gorm:"not null"`
	UpdatedAt  time.Time      `gorm:"not null"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	User       *User          `gorm:"foreignkey:UserID" json:"user,omitempty"`
}
