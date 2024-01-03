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
