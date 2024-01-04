package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Vehicle struct {
	ID                uint      `gorm:"primaryKey;autoIncrement"`
	UUID              uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();uniqueIndex"`
	OwnerID           uint
	CategoryID        uint
	LocationID        uint
	Status            string `gorm:"type:varchar(100);not null"`
	PricePerHour      float64
	PricePerDay       float64
	CreatedAt         time.Time          `gorm:"not null"`
	UpdatedAt         time.Time          `gorm:"not null"`
	DeletedAt         gorm.DeletedAt     `gorm:"index"`
	Category          VehicleCategory    `gorm:"foreignkey:CategoryID"`
	Location          Location           `gorm:"foreignkey:LocationID"`
	Rentals           []Rental           `gorm:"foreignkey:VehicleID"`
	VehicleInsurances []VehicleInsurance `gorm:"foreignkey:VehicleID"`
}

type VehicleResponse struct {
	ID           uuid.UUID `json:"id,omitempty"`
	OwnerID      uint      `json:"owner_id,omitempty"`
	CategoryID   uint      `json:"category_id,omitempty"`
	LocationID   uint      `json:"location_id,omitempty"`
	Status       string    `json:"status,omitempty"`
	PricePerHour float64   `json:"price_per_hour,omitempty"`
	PricePerDay  float64   `json:"price_per_day,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
}
