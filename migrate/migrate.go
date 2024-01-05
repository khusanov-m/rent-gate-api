package main

import (
	"fmt"
	"log"

	"github.com/khusanov-m/rent-gate-api/initializers"
	"github.com/khusanov-m/rent-gate-api/models"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	err := initializers.DB.AutoMigrate(
		&models.User{},
		&models.Role{},
		&models.VehicleCategory{},
		&models.Vehicle{},
		&models.Location{},
		&models.Rental{},
		&models.Subscription{},
		&models.InsuranceOption{},
		&models.VehicleInsurance{},
		&models.LoyaltyProgram{},
		&models.Post{},
	)
	if err != nil {
		panic("failed to auto-migrate")
	}
	fmt.Println("? Migration complete")
}
