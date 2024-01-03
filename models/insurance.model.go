package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type InsuranceOption struct {
	ID                uint      `gorm:"primaryKey;autoIncrement"`
	UUID              uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();uniqueIndex"`
	Name              string    `gorm:"type:varchar(255);not null"`
	Description       string    `gorm:"type:text"`
	Price             float64
	CreatedAt         time.Time      `gorm:"not null"`
	UpdatedAt         time.Time      `gorm:"not null"`
	DeletedAt         gorm.DeletedAt `gorm:"index"`
	VehicleInsurances []VehicleInsurance
}

type VehicleInsurance struct {
	ID                uint      `gorm:"primaryKey;autoIncrement"`
	UUID              uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();uniqueIndex"`
	VehicleID         uint
	InsuranceOptionID uint
	CreatedAt         time.Time       `gorm:"not null"`
	UpdatedAt         time.Time       `gorm:"not null"`
	DeletedAt         gorm.DeletedAt  `gorm:"index"`
	Vehicle           Vehicle         `gorm:"foreignkey:VehicleID"`
	InsuranceOption   InsuranceOption `gorm:"foreignkey:InsuranceOptionID"`
}
