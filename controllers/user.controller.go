package controllers

// CAN PRELOAD: Preload all by default

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/khusanov-m/rent-gate-api/models"
	"github.com/khusanov-m/rent-gate-api/utils"
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
	userResponse := utils.MapUserToUserResponse(&currentUser)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": userResponse}})
}
