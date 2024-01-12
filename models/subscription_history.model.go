package models

import (
	"github.com/google/uuid"
	"time"
)

type SubscriptionHistory struct {
	ID                 uint      `gorm:"primaryKey;autoIncrement"`
	UUID               uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();uniqueIndex"`
	UserID             uint      `gorm:"not null"`
	SubscriptionTypeID uuid.UUID `gorm:"not null"`
	SubscriptionPrice  float64   `gorm:"not null"`
	Vehicles           []Vehicle `gorm:"many2many:subscription_history_vehicles;"`
	StartDate          time.Time `gorm:"not null"`
	EndDate            time.Time `gorm:"not null"`

	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
	DeletedAt time.Time `gorm:"index"`
}

type SubscriptionHistoryInput struct {
	UserID             uint      `json:"user_id,omitempty"`
	SubscriptionTypeID uuid.UUID `json:"subscription_type_id,omitempty"`
	SubscriptionPrice  float64   `json:"subscription_price,omitempty"`
	Vehicles           []Vehicle `json:"vehicles,omitempty"`
	StartDate          time.Time `json:"start_date,omitempty"`
	EndDate            time.Time `json:"end_date,omitempty"`
}

type SubscriptionHistoryResponse struct {
	ID                uuid.UUID `json:"user_id,omitempty"`
	SubscriptionPrice float64   `json:"subscription_price,omitempty"`
	Vehicles          []Vehicle `json:"vehicles,omitempty"`
	StartDate         time.Time `json:"start_date,omitempty"`
	EndDate           time.Time `json:"end_date,omitempty"`
}
