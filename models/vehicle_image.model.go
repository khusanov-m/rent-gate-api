package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type VehicleImage struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	UUID      uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();uniqueIndex"`
	VehicleID uint      `json:"vehicle_id"`
	ImageURL  string    `json:"image_url"`

	CreatedAt time.Time      `gorm:"not null"`
	UpdatedAt time.Time      `gorm:"not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type VehicleImageResponse struct {
	ID        uuid.UUID `json:"id,omitempty"`
	VehicleID uint      `json:"vehicle_id,omitempty"`
	ImageURL  string    `json:"image_url,omitempty"`
}
