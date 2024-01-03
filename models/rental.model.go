package models

import (
	"time"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

type Rental struct {
	ID         uint      `gorm:"primaryKey;autoIncrement"`
	UUID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();uniqueIndex"`
	UserID     uint
	VehicleID  uint
	StartDate  time.Time
	EndDate    time.Time
	TotalPrice float64
	CreatedAt  time.Time      `gorm:"not null"`
	UpdatedAt  time.Time      `gorm:"not null"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	User       User           `gorm:"foreignkey:UserID"`
	Vehicle    Vehicle        `gorm:"foreignkey:VehicleID"`
}
