package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Vehicle struct {
	ID                 uint      `gorm:"primaryKey;autoIncrement"`
	UUID               uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();uniqueIndex"`
	AvailabilityStatus bool      `gorm:"not null"`
	DriverOption       string    `gorm:"type:varchar(100);not null"` // WithDriver, WithoutDriver, Both
	PricePerHour       float64   `gorm:"not null"`
	PricePerDay        float64   `gorm:"not null"`
	NumberOfSeats      uint16    `gorm:"not null"`
	LuggageCapacity    float32   `gorm:"not null"`
	VehicleType        string    `gorm:"type:varchar(100);not null"` // Car, Motorbike, Bicycle, Boat, Plane
	PowerType          string    `gorm:"type:varchar(100);not null"` // Petrol, Diesel, Electric, Hybrid
	ImageList          []string  `gorm:"type:varchar(100);not null"`

	CreatedAt time.Time      `gorm:"not null"`
	UpdatedAt time.Time      `gorm:"not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	OwnerType          string `gorm:"type:varchar(100);not null"` // User, Company
	OwnerID            uint
	Location           *Location `gorm:"foreignkey:Identification"`
	SubscriptionTierID *uint
}

type CreateVehicleInput struct {
	Status             string   `json:"status" binding:"required"`
	AvailabilityStatus bool     `json:"availability_status" binding:"required"`
	DriverOption       string   `json:"driver_option" binding:"required"`
	PricePerHour       float64  `json:"price_per_hour" binding:"required"`
	PricePerDay        float64  `json:"price_per_day" binding:"required"`
	NumberOfSeats      uint16   `json:"number_of_seats" binding:"required"`
	LuggageCapacity    float32  `json:"luggage_capacity" binding:"required"`
	VehicleType        string   `json:"vehicle_type" binding:"required"`
	PowerType          string   `json:"power_type" binding:"required"`
	ImageList          []string `json:"image_list" binding:"required"`

	OwnerType          string `json:"owner_type" binding:"required"`
	OwnerID            uint   `json:"owner_id" binding:"required"`
	SubscriptionTierID *uint  `json:"subscription_tier_id"`
}

type VehicleResponse struct {
	ID                 uuid.UUID `json:"id,omitempty"`
	AvailabilityStatus bool      `json:"availability_status,omitempty"`
	DriverOption       string    `json:"driver_option,omitempty"`
	PricePerHour       float64   `json:"price_per_hour,omitempty"`
	PricePerDay        float64   `json:"price_per_day,omitempty"`
	NumberOfSeats      uint16    `json:"number_of_seats,omitempty"`
	LuggageCapacity    float32   `json:"luggage_capacity,omitempty"`
	VehicleType        string    `json:"vehicle_type,omitempty"`
	PowerType          string    `json:"power_type,omitempty"`
	ImageList          []string  `json:"image_list,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	OwnerType          string    `json:"owner_type,omitempty"`
	OwnerID            uint      `json:"owner_id,omitempty"`
	Location           *Location `json:"location,omitempty"`
	SubscriptionTierID *uint     `json:"subscription_tier_id,omitempty"`
}
