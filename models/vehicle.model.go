package models

import (
	"gorm.io/gorm"
	"time"

	"github.com/google/uuid"
)

type Vehicle struct {
	ID              uint      `gorm:"primaryKey;autoIncrement"`
	UUID            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();uniqueIndex"`
	IsAvailable     bool
	DriverOption    string  `gorm:"type:varchar(100)"` // WithDriver, WithoutDriver, Both
	PricePerHour    float64 `gorm:"not null"`
	PricePerDay     float64 `gorm:"not null"`
	Currency        string  `gorm:"type:varchar(100);not null"` // currency code
	NumberOfSeats   uint16  // counts adults
	LuggageCapacity float32 // in kg
	VehicleType     string  `gorm:"type:varchar(100)"` // Car, Motorbike, Bicycle, Boat, Plane
	PowerType       string  `gorm:"type:varchar(100)"` // Petrol, Diesel, Electric, Hybrid
	Make            string  `gorm:"type:varchar(100)"` // Manufacturer: Chevrolet, Tesla, Ford, etc.
	Model           string  `gorm:"type:varchar(100)"` // Car name: Model Y, Tahoe, Impala, Gentra
	Color           string  `gorm:"type:varchar(100)"`

	OwnerType          string `gorm:"type:varchar(100);not null"` // User, Company
	OwnerID            uint
	Location           *Location         `gorm:"foreignkey:Identification"`
	InSubscriptionType *SubscriptionType `gorm:"foreignkey:VehicleID"`
	Images             []VehicleImage    `gorm:"foreignkey:VehicleID"`

	CreatedAt time.Time      `gorm:"not null"`
	UpdatedAt time.Time      `gorm:"not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type CreateVehicleInput struct {
	IsAvailable     bool           `json:"is_available"`
	DriverOption    string         `json:"driver_option" binding:"required"`
	PricePerHour    float64        `json:"price_per_hour" binding:"required"`
	PricePerDay     float64        `json:"price_per_day" binding:"required"`
	Currency        string         `json:"currency" binding:"required"`
	NumberOfSeats   uint16         `json:"number_of_seats" binding:"required"`
	LuggageCapacity float32        `json:"luggage_capacity" binding:"required"`
	VehicleType     string         `json:"vehicle_type" binding:"required"`
	PowerType       string         `json:"power_type" binding:"required"`
	Images          []VehicleImage `json:"images" binding:"required"`
	Make            string         `json:"make"`
	Model           string         `json:"model"`
	Color           string         `json:"color"`
}

type UpdateVehicleInput struct {
	IsAvailable     bool           `json:"is_available"`
	DriverOption    string         `json:"driver_option"`
	PricePerHour    float64        `json:"price_per_hour"`
	PricePerDay     float64        `json:"price_per_day"`
	Currency        string         `json:"currency"`
	NumberOfSeats   uint16         `json:"number_of_seats"`
	LuggageCapacity float32        `json:"luggage_capacity"`
	VehicleType     string         `json:"vehicle_type"`
	PowerType       string         `json:"power_type"`
	Images          []VehicleImage `json:"images"`
	Make            string         `json:"make"`
	Model           string         `json:"model"`
	Color           string         `json:"color"`
}

type VehicleResponse struct {
	ID              uuid.UUID `json:"id,omitempty"`
	IsAvailable     bool      `json:"is_available,omitempty"`
	DriverOption    string    `json:"driver_option,omitempty"`
	PricePerHour    float64   `json:"price_per_hour,omitempty"`
	PricePerDay     float64   `json:"price_per_day,omitempty"`
	Currency        string    `json:"currency,omitempty"`
	NumberOfSeats   uint16    `json:"number_of_seats,omitempty"`
	LuggageCapacity float32   `json:"luggage_capacity,omitempty"`
	VehicleType     string    `json:"vehicle_type,omitempty"`
	PowerType       string    `json:"power_type,omitempty"`
	Make            string    `json:"make,omitempty"`
	Model           string    `json:"model,omitempty"`
	Color           string    `json:"color,omitempty"`

	OwnerType          string                 `json:"owner_type"`
	OwnerID            uint                   `json:"owner_id"`
	Location           *Location              `json:"location"`
	InSubscriptionType *SubscriptionType      `json:"in_subscription_type"`
	Images             []VehicleImageResponse `json:"images"`

	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
