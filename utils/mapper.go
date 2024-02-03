package utils

import (
	"github.com/khusanov-m/rent-gate-api/models"
	"math"
)

func MapUsersToUsersResponse(users *[]models.User) []models.UserResponse {
	usersResponse := make([]models.UserResponse, len(*users))
	for i, user := range *users {
		usersResponse[i] = MapUserToUserResponse(&user)
	}
	return usersResponse
}
func MapUserToUserResponse(user *models.User) models.UserResponse {
	// vehiclesResponse := MapVehiclesToVehicleResponses(&user.Vehicles)
	// postsResponse := MapPostsToPostResponses(&user.Posts)

	return models.UserResponse{
		ID:        user.UUID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		Photo:     user.Photo,
		Provider:  user.Provider,
		Verified:  user.Verified,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
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
	vehicleImagesResponse := MapVehicleImagesToVehicleImagesResponse(&vehicle.Images)

	return models.VehicleResponse{
		ID:              vehicle.UUID,
		PricePerHour:    vehicle.PricePerHour,
		PricePerDay:     vehicle.PricePerDay,
		Location:        vehicle.Location,
		IsAvailable:     vehicle.IsAvailable,
		DriverOption:    vehicle.DriverOption,
		NumberOfSeats:   vehicle.NumberOfSeats,
		LuggageCapacity: vehicle.LuggageCapacity,
		VehicleType:     vehicle.VehicleType,
		PowerType:       vehicle.PowerType,
		Currency:        vehicle.Currency,
		OwnerType:       vehicle.OwnerType,
		OwnerID:         vehicle.OwnerID,
		Images:          *vehicleImagesResponse,
		Model:           vehicle.Model,
		Make:            vehicle.Make,
		Color:           vehicle.Color,
		CreatedAt:       vehicle.CreatedAt,
		UpdatedAt:       vehicle.UpdatedAt,
	}
}

//func MapPostsToPostResponses(posts *[]models.Post) []models.PostResponse {
//	postsResponses := make([]models.PostResponse, len(*posts))
//	for i, post := range *posts {
//		postsResponses[i] = MapPostToPostResponse(&post)
//	}
//	return postsResponses
//}
//func MapPostToPostResponse(post *models.Post) models.PostResponse {
//	if post.User != nil {
//		userResponse := MapUserToPreloadUserResponse(post.User)
//
//		return models.PostResponse{
//			ID:      post.UUID,
//			Title:   post.Title,
//			Content: post.Content,
//			Image:   post.Image,
//			User:    userResponse,
//		}
//	}
//
//	return models.PostResponse{
//		ID:      post.UUID,
//		Title:   post.Title,
//		Content: post.Content,
//		Image:   post.Image,
//	}
//}

//func MapUserToPreloadUserResponse(user *models.User) *models.UserResponse {
//	return &models.UserResponse{
//		ID:       user.UUID,
//		Name:     user.Name,
//		Email:    user.Email,
//		Role:     user.Role,
//		Photo:    user.Photo,
//		Provider: user.Provider,
//		Verified: user.Verified,
//	}
//}

func MapVehicleImagesToVehicleImagesResponse(vehicleImages *[]models.VehicleImage) *[]models.VehicleImageResponse {
	vehicleImagesResponse := make([]models.VehicleImageResponse, len(*vehicleImages))
	for i, image := range *vehicleImages {
		vehicleImagesResponse[i] = *MapVehicleImageToVehicleImageResponse(&image)
	}
	return &vehicleImagesResponse
}
func MapVehicleImageToVehicleImageResponse(vehicleImage *models.VehicleImage) *models.VehicleImageResponse {
	return &models.VehicleImageResponse{
		ID:        vehicleImage.UUID,
		ImageURL:  vehicleImage.ImageURL,
		VehicleID: vehicleImage.VehicleID,
	}
}
func MapPaymentsToPaymentsResponse(payments *[]models.Payment) []models.PaymentResponse {
	paymentsRes := make([]models.PaymentResponse, len(*payments))
	for i, payment := range *payments {
		paymentsRes[i] = MapPaymentToPaymentResponse(&payment)
	}
	return paymentsRes
}
func MapPaymentToPaymentResponse(payment *models.Payment) models.PaymentResponse {
	return models.PaymentResponse{
		ID:             payment.UUID,
		UserID:         payment.UserID,
		Amount:         math.Ceil(payment.Amount*100) / 100,
		PaymentStatus:  payment.PaymentStatus,
		PaymentType:    payment.PaymentType,
		PaymentFor:     payment.PaymentFor,
		PaymentDetails: payment.PaymentDetails,
		CreatedAt:      payment.CreatedAt,
		UpdatedAt:      payment.UpdatedAt,
	}
}
