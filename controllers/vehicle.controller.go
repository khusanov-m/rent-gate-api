package controllers

// CAN PRELOAD: User

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/khusanov-m/rent-gate-api/models"
	"github.com/khusanov-m/rent-gate-api/utils"
	"gorm.io/gorm"
)

type VehicleController struct {
	DB *gorm.DB
}

func NewVehicleController(DB *gorm.DB) VehicleController {
	return VehicleController{DB}
}

var vehicleAallowedEntities utils.PreloadEntities = utils.PreloadEntities{
	"Rentals":           true,
	"VehicleCategory":   true,
	"Location":          true,
	"VehicleInsurances": true,
}

// [...] Create Vehicle Handler
func (vc *VehicleController) CreateVehicle(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	var payload *models.CreateVehicleInput

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	now := time.Now()
	newVehicle := models.Vehicle{
		UserID:       currentUser.ID,
		Status:       payload.Status,
		PricePerHour: payload.PricePerHour,
		PricePerDay:  payload.PricePerDay,
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	result := vc.DB.Create(&newVehicle)
	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key") {
		ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Vehicle with that title already exists"})
		return
	} else if result.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newVehicle})
}

func (vc *VehicleController) GetVehicles(ctx *gin.Context) {
	var vehicles []models.Vehicle

	query := utils.ApplyDynamicPreloading(vc.DB, ctx, vehicleAallowedEntities)

	if err := query.Find(&vehicles).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"vehicles": vehicles}})
}

func (vc *VehicleController) GetVehicleByID(ctx *gin.Context) {
	var vehicle models.Vehicle
	id := ctx.Param("id")

	query := utils.ApplyDynamicPreloading(vc.DB, ctx, vehicleAallowedEntities)

	if err := query.Where("uuid = ?", id).Preload("User").First(&vehicle).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	vehicleResponse := utils.MapVehicleToVehicleResponse(&vehicle)
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"vehicle": vehicleResponse}})
}
