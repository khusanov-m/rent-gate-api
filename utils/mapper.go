package utils

import (
	"github.com/khusanov-m/rent-gate-api/models"
)

func MapUserToUserResponse(user *models.User) models.UserResponse {
	// vehiclesResponse := MapVehiclesToVehicleResponses(&user.Vehicles)
	// postsResponse := MapPostsToPostResponses(&user.Posts)

	return models.UserResponse{
		ID:        user.UUID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		PhotoUrl:  user.PhotoUrl,
		Provider:  user.Provider,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Verified:  user.Verified,
		// Vehicles:       vehiclesResponse,
		// Subscription:   user.Subscription,
		// LoyaltyProgram: user.LoyaltyProgram,
		// Posts: postsResponse,
	}
}

func MapVehiclesToVehicleResponses(vehicles *[]models.Vehicle) []models.VehicleResponse {
	vehiclesResponse := make([]models.VehicleResponse, len(*vehicles))
	for i, vehicle := range *vehicles {
		vehiclesResponse[i] = MapVehicleToVehicleResponse(&vehicle)
	}
	return vehiclesResponse
}
func MapVehicleToVehicleResponse(vehicle *models.Vehicle) models.VehicleResponse {
	return models.VehicleResponse{
		ID: vehicle.UUID,
		// UserID:            vehicle.UserID,
		// Status:            vehicle.Status,
		PricePerHour: vehicle.PricePerHour,
		PricePerDay:  vehicle.PricePerDay,
		CreatedAt:    vehicle.CreatedAt,
		UpdatedAt:    vehicle.UpdatedAt,
		// Category:          vehicle.Category,
		Location: vehicle.Location,
		// Rentals:           vehicle.Rentals,
		// VehicleInsurances: vehicle.VehicleInsurances,
	}
}

func MapRentalsToRentalResponses(rentals *[]models.Rental) []models.RentalResponse {
	rentalsResponse := make([]models.RentalResponse, len(*rentals))
	for i, rental := range *rentals {
		rentalsResponse[i] = MapRentalToRentalResponse(&rental)
	}
	return rentalsResponse
}
func MapRentalToRentalResponse(rental *models.Rental) models.RentalResponse {
	return models.RentalResponse{
		ID:         rental.UUID,
		VehicleID:  rental.VehicleID,
		StartDate:  rental.StartDate,
		EndDate:    rental.EndDate,
		TotalPrice: rental.TotalPrice,
		CreatedAt:  rental.CreatedAt,
		UpdatedAt:  rental.UpdatedAt,
	}
}

func MapPostsToPostResponses(posts *[]models.Post) []models.PostResponse {
	postsResponses := make([]models.PostResponse, len(*posts))
	for i, post := range *posts {
		postsResponses[i] = MapPostToPostResponse(&post)
	}
	return postsResponses
}
func MapPostToPostResponse(post *models.Post) models.PostResponse {
	if post.User != nil {
		userResponse := MapUserToPreloadUserResponse(post.User)

		return models.PostResponse{
			ID:      post.UUID,
			Title:   post.Title,
			Content: post.Content,
			Image:   post.Image,
			User:    userResponse,
		}
	}

	return models.PostResponse{
		ID:      post.UUID,
		Title:   post.Title,
		Content: post.Content,
		Image:   post.Image,
	}
}

func MapUserToPreloadUserResponse(user *models.User) *models.UserResponse {
	return &models.UserResponse{
		ID:       user.UUID,
		Name:     user.Name,
		Email:    user.Email,
		Role:     user.Role,
		PhotoUrl: user.PhotoUrl,
		Provider: user.Provider,
		Verified: user.Verified,
	}
}
