package models

import (
	"time"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

// RentalSchema
type Rental struct {
	ID         uint      `gorm:"primaryKey;autoIncrement"`
	UUID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();uniqueIndex"`
	VehicleID  uint
	StartDate  time.Time
	EndDate    time.Time
	TotalPrice float64
	CreatedAt  time.Time      `gorm:"not null"`
	UpdatedAt  time.Time      `gorm:"not null"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

type RentalResponse struct {
	ID         uuid.UUID `json:"id,omitempty"`
	VehicleID  uint      `json:"vehicle_id,omitempty"`
	StartDate  time.Time `json:"start_date,omitempty"`
	EndDate    time.Time `json:"end_date,omitempty"`
	TotalPrice float64   `json:"total_price,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
}
