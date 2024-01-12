package models

import (
	"github.com/google/uuid"
	"time"
)

type SubscriptionType struct {
	ID    uint      `gorm:"primaryKey;autoIncrement"`
	UUID  uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();uniqueIndex"`
	Type  string    `gorm:"type:varchar(255);uniqueIndex;not null"` // no duplicates
	Price float64   `gorm:"not null"`

	VehicleID uint

	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
	DeletedAt time.Time `gorm:"index"`
}

type SubscriptionTypeInput struct {
	Type  string  `json:"type,omitempty"`
	Price float64 `json:"price,omitempty"`
}

type SubscriptionTypeResponse struct {
	ID    uuid.UUID `json:"id,omitempty"`
	Type  string    `json:"type,omitempty"`
	Price float64   `json:"price,omitempty"`
}
