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
		&models.Location{},
		&models.Vehicle{},
		//&models.Company{},
		&models.LoyaltyAccount{},
		&models.Payment{},
		&models.RentPaymentHistory{},
		&models.Rental{},
		&models.Subscription{},
		&models.SubscriptionHistory{},
		&models.SubscriptionType{},
	)
	if err != nil {
		panic("failed to auto-migrate")
	}
	fmt.Println("? Migration complete")
}
