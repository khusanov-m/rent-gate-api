package models

import (
	"github.com/google/uuid"
	"time"
)

type RentPaymentHistory struct {
	ID           uint      `gorm:"primaryKey;autoIncrement"`
	UUID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();uniqueIndex"`
	VehicleID    uint      `gorm:"not null"`
	UserID       uint      `gorm:"not null"`
	TotalAmount  float64   `gorm:"not null"`
	PricePerHour float64   `gorm:"not null"`
	PricePerDay  float64   `gorm:"not null"`
	Duration     uint      `gorm:"not null"`

	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
	DeletedAt time.Time `gorm:"index"`
}

type RentPaymentHistoryInput struct {
	VehicleID    uuid.UUID `json:"vehicle_id,omitempty"`
	TotalAmount  float64   `json:"total_amount,omitempty"`
	PricePerHour float64   `json:"price_per_hour,omitempty"`
	PricePerDay  float64   `json:"price_per_day,omitempty"`
	Duration     uint      `json:"duration,omitempty"`
}

type RentPaymentHistoryResponse struct {
	ID           uuid.UUID `json:"id,omitempty"`
	TotalAmount  float64   `json:"total_amount,omitempty"`
	PricePerHour float64   `json:"price_per_hour,omitempty"`
	PricePerDay  float64   `json:"price_per_day,omitempty"`
	Duration     uint      `json:"duration,omitempty"`
}
