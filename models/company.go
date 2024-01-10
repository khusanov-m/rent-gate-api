package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Company struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id,omitempty"`
	UUID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();uniqueIndex" json:"uuid,omitempty"`
	Name        string    `gorm:"type:varchar(255);not null" json:"name,omitempty"`
	Description string    `gorm:"type:text" json:"description,omitempty"`
	Email       string    `gorm:"type:varchar(255);not null" json:"email,omitempty"`
	Phone       string    `gorm:"type:varchar(255);not null" json:"phone,omitempty"`
	Address     string    `gorm:"type:varchar(255);not null" json:"address,omitempty"`

	CreatedAt time.Time      `gorm:"not null"`
	UpdatedAt time.Time      `gorm:"not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Vehicles []*Vehicle `gorm:"foreignKey:OwnerID" json:"vehicles,omitempty"`
}

type CompanyResponse struct {
	ID          uuid.UUID `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	Email       string    `json:"email,omitempty"`
	Phone       string    `json:"phone,omitempty"`
	Address     string    `json:"address,omitempty"`

	Vehicles []*VehicleResponse `json:"vehicles,omitempty"`
}
