package utils

import (
	"github.com/khusanov-m/rent-gate-api/models"
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
	vehicleImagesResponse := MapVehicleImagesToVehicleImagesResponse(&vehicle.Images)

	return models.VehicleResponse{
		ID:              vehicle.UUID,
		PricePerHour:    vehicle.PricePerHour,
		PricePerDay:     vehicle.PricePerDay,
		CreatedAt:       vehicle.CreatedAt,
		UpdatedAt:       vehicle.UpdatedAt,
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
