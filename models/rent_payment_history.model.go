package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type RentPaymentHistory struct {
	ID           uint      `gorm:"primaryKey;autoIncrement"`
	UUID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();uniqueIndex"`
	VehicleID    uuid.UUID `gorm:"not null"`
	PaymentID    uuid.UUID `gorm:"not null"`
	UserID       uint      `gorm:"not null"`
	TotalAmount  float64   `gorm:"not null"`
	PricePerHour float64   `gorm:"not null"`
	PricePerDay  float64   `gorm:"not null"`
	Duration     uint      `gorm:""`
	Status       string    `gorm:"type:varchar(100)"`

	CreatedAt time.Time      `gorm:"not null"`
	UpdatedAt time.Time      `gorm:"not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type RentPaymentHistoryResponse struct {
	ID           uuid.UUID `json:"id,omitempty"`
	UserID       uint      `json:"user_id,omitempty"`
	VehicleID    uuid.UUID `json:"vehicle_id,omitempty"`
	PaymentID    uuid.UUID `json:"payment_id,omitempty"`
	TotalAmount  float64   `json:"total_amount,omitempty"`
	PricePerHour float64   `json:"price_per_hour,omitempty"`
	PricePerDay  float64   `json:"price_per_day,omitempty"`
	Duration     uint      `json:"duration,omitempty"`
	Status       string    `json:"status,omitempty"`
}
