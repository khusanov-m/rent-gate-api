package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Location struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	UUID      uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();uniqueIndex"`
	Address   string    `gorm:"type:varchar(255);not null"`
	Latitude  float64   `gorm:"not null"`
	Longitude float64   `gorm:"not null"`
	Country   string    `gorm:"type:varchar(255)"`
	City      string    `gorm:"type:varchar(255)"`
	District  string    `gorm:"type:varchar(255)"`
	ZipCode   string    `gorm:"type:varchar(255)"`

	CreatedAt time.Time      `gorm:"not null"`
	UpdatedAt time.Time      `gorm:"not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	VehicleID uint
}
