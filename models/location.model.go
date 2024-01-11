package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Location struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	UUID      uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();uniqueIndex"`
	Latitude  float64   `gorm:"not null"`
	Longitude float64   `gorm:"not null"`
	Address   string    `gorm:"type:varchar(255);not null"`
	Country   string    `gorm:"type:varchar(255);not null"`
	City      string    `gorm:"type:varchar(255);not null"`
	District  string    `gorm:"type:varchar(255);not null"`
	ZipCode   string    `gorm:"type:varchar(255);not null"`

	CreatedAt time.Time      `gorm:"not null"`
	UpdatedAt time.Time      `gorm:"not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Identification uint
}
