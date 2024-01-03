package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	ID        uint           `gorm:"primaryKey;autoIncrement"`
	UUID      uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();uniqueIndex"`
	Name      string         `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time      `gorm:"not null"`
	UpdatedAt time.Time      `gorm:"not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	// Users     []User
}
