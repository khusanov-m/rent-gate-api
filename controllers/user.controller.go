package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/khusanov-m/rent-gate-api/models"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserController(DB *gorm.DB) UserController {
	return UserController{DB}
}

func (uc *UserController) GetMe(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)

	userResponse := &models.UserResponse{
		ID:             currentUser.UUID,
		Name:           currentUser.Name,
		Email:          currentUser.Email,
		Photo:          currentUser.Photo,
		Role:           currentUser.Role,
		Provider:       currentUser.Provider,
		CreatedAt:      currentUser.CreatedAt,
		UpdatedAt:      currentUser.UpdatedAt,
		Verfied:        currentUser.Verified,
		Rentals:        currentUser.Rentals,
		Vehicles:       currentUser.Vehicles,
		Subscription:   currentUser.Subscription,
		LoyaltyProgram: currentUser.LoyaltyProgram,
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": userResponse}})
}
