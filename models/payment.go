package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Payment struct {
	ID            uint      `gorm:"primaryKey;autoIncrement"`
	UUID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();uniqueIndex"`
	UserID        uint      `gorm:"not null"`
	Amount        float64   `gorm:"not null"`
	PaymentStatus string    `gorm:"type:varchar(255);not null"`
	PaymentType   string    `gorm:"type:varchar(255);not null"`

	PaymentFor     string    `gorm:"type:varchar(255);not null"` // rent, subscription
	PaymentDetails uuid.UUID `gorm:"type:uuid"`                  // not connected by FK, but can be joined via UUID manually, UUID of vehicle

	CreatedAt time.Time      `gorm:"not null"`
	UpdatedAt time.Time      `gorm:"not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type PaymentInput struct {
	TotalHours              uint    `json:"total_hours"`
	AddonsWithDiscountPrice float64 `json:"addons_with_discount_price"`
	PaymentType             string  `json:"payment_type"`
}

type PaymentResponse struct {
	ID            uuid.UUID `json:"id,omitempty"`
	UserID        uint      `json:"user_id,omitempty"`
	Amount        float64   `json:"amount,omitempty"`
	PaymentStatus string    `json:"payment_status,omitempty"`
	PaymentType   string    `json:"payment_type,omitempty"`

	PaymentFor     string    `json:"payment_for,omitempty"`
	PaymentDetails uuid.UUID `json:"payment_details,omitempty"` // not connected by FK, but can be joined via UUID manually

	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
