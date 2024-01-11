package models

import "github.com/google/uuid"

type SubscriptionType struct {
	ID    uint      `gorm:"primaryKey;autoIncrement"`
	UUID  uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();uniqueIndex"`
	Type  string    `gorm:"type:varchar(255);not null"`
	Price float64   `gorm:"not null"`
}
