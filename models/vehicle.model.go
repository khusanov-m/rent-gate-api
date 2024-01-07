package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Vehicle struct {
	ID           uint      `gorm:"primaryKey;autoIncrement"`
	UUID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();uniqueIndex"`
	UserID       uint
	Status       string `gorm:"type:varchar(100);not null"`
	PricePerHour float64
	PricePerDay  float64
	CreatedAt    time.Time      `gorm:"not null"`
	UpdatedAt    time.Time      `gorm:"not null"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`

	Category          *VehicleCategory   `gorm:"foreignkey:VehicleID"`
	Location          *Location          `gorm:"foreignkey:VehicleID"`
	Rentals           []Rental           `gorm:"foreignkey:VehicleID"`
	VehicleInsurances []VehicleInsurance `gorm:"foreignkey:VehicleID"`
}

type CreateVehicleInput struct {
	Status       string  `json:"status" binding:"required"`
	PricePerHour float64 `json:"price_per_hour" binding:"required"`
	PricePerDay  float64 `json:"price_per_day" binding:"required"`
}

type VehicleResponse struct {
	ID                uuid.UUID          `json:"id,omitempty"`
	UserID            uint               `json:"owner_id,omitempty"`
	CategoryID        uint               `json:"category_id,omitempty"`
	LocationID        uint               `json:"location_id,omitempty"`
	Status            string             `json:"status,omitempty"`
	PricePerHour      float64            `json:"price_per_hour,omitempty"`
	PricePerDay       float64            `json:"price_per_day,omitempty"`
	CreatedAt         time.Time          `json:"created_at,omitempty"`
	UpdatedAt         time.Time          `json:"updated_at,omitempty"`
	Category          *VehicleCategory   `json:"category,omitempty"`
	Location          *Location          `json:"location,omitempty"`
	Rentals           []Rental           `json:"rentals,omitempty"`
	VehicleInsurances []VehicleInsurance `json:"vehicle_insurances,omitempty"`
}
