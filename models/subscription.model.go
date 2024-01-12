package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Subscription struct {
	ID                 uint      `gorm:"primaryKey;autoIncrement"`
	UUID               uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();uniqueIndex"`
	UserID             uint      `gorm:"not null"`
	SubscriptionTypeID uuid.UUID `gorm:"not null"`
	StartDate          time.Time `gorm:"not null"`
	EndDate            time.Time `gorm:"not null"`
	TotalPrice         float64   `gorm:"not null"`

	Vehicles []Vehicle `gorm:"many2many:subscription_vehicles;"`

	CreatedAt time.Time      `gorm:"not null"`
	UpdatedAt time.Time      `gorm:"not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	//User *User `gorm:"foreignkey:UserID" json:"user,omitempty"`
}

type SubscriptionResponse struct {
	ID                 uuid.UUID `json:"id,omitempty"`
	UserID             uint      `json:"user_id,omitempty"`
	SubscriptionTypeID uuid.UUID `json:"subscription_type_id,omitempty"`
	StartDate          time.Time `json:"start_date,omitempty"`
	EndDate            time.Time `json:"end_date,omitempty"`
	TotalPrice         float64   `json:"total_price,omitempty"`

	Vehicles []Vehicle `json:"vehicles,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

/*
var subscription Subscription
db.Preload("VehicleList").Find(&subscription, "id = ?", someSubscriptionID)
*/
