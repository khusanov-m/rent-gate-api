package models

import (
	"github.com/google/uuid"
	"time"
)

type Payment struct {
	ID             uint      `gorm:"primaryKey;autoIncrement"`
	UUID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();uniqueIndex"`
	UserID         uint      `gorm:"not null"`
	Amount         float64   `gorm:"not null"`
	PaymentStatus  string    `gorm:"type:varchar(255);not null"`
	PaymentFor     string    `gorm:"type:varchar(255);not null"`
	PaymentDetails uuid.UUID `gorm:"not null"` // not connected by FK, but can be joined via UUID manually
	PaymentDate    time.Time `gorm:"not null"`
	PaymentType    string    `gorm:"type:varchar(255)"`
}
