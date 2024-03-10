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
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "user": userResponse})
}

func (uc *UserController) GetAllUsers(ctx *gin.Context) {
	var users []models.User

	if err := uc.DB.Find(&users).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	usersResponse := utils.MapUsersToUsersResponse(&users)
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "users": usersResponse})
}
