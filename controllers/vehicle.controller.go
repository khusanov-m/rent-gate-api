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

var vehicleAllowedEntities utils.PreloadEntities = utils.PreloadEntities{
	"Location": true,
	"Images":   true,
	"User":     true,
}

// [...] Create Vehicle Handler
func (vc *VehicleController) CreateVehicle(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	var payload *models.CreateVehicleInput

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	newVehicle := models.Vehicle{
		IsAvailable:     payload.IsAvailable,
		DriverOption:    payload.DriverOption,
		NumberOfSeats:   payload.NumberOfSeats,
		LuggageCapacity: payload.LuggageCapacity,
		VehicleType:     payload.VehicleType,
		PowerType:       payload.PowerType,
		PricePerHour:    payload.PricePerHour,
		PricePerDay:     payload.PricePerDay,
		Currency:        payload.Currency,
		Model:           payload.Model,
		Make:            payload.Make,
		Color:           payload.Color,

		OwnerType: currentUser.Role,
		OwnerID:   currentUser.ID,

		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	result := vc.DB.Create(&newVehicle)
	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key") {
		ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Vehicle with that id already exists"})
		return
	} else if result.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	for _, image := range payload.Images {
		vehicleImage := models.VehicleImage{
			VehicleID: newVehicle.ID,
			ImageURL:  image.ImageURL,
		}
		vc.DB.Create(&vehicleImage)
	}

	newVehicleResponse := utils.MapVehicleToVehicleResponse(&newVehicle)
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newVehicleResponse})
}

// [...] Update Vehicle Handler
func (vc *VehicleController) UpdateVehicle(ctx *gin.Context) {
	postId := ctx.Param("vehicleId")
	currentUser := ctx.MustGet("currentUser").(models.User)

	var payload *models.UpdateVehicleInput
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var vehicle models.Vehicle
	result := vc.DB.First(&vehicle, "uuid = ?", postId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
		return
	}

	if vehicle.OwnerID != currentUser.ID {
		ctx.JSON(http.StatusForbidden, gin.H{"status": "fail", "message": "You are not authorized to update this post"})
		return
	}

	var updatedVehicle models.Vehicle
	vc.DB.First(&updatedVehicle, "uuid = ?", postId)

	vehicleToUpdate := models.Vehicle{
		IsAvailable:     payload.IsAvailable,
		DriverOption:    payload.DriverOption,
		NumberOfSeats:   payload.NumberOfSeats,
		Color:           payload.Color,
		Make:            payload.Make,
		Model:           payload.Model,
		Currency:        payload.Currency,
		PowerType:       payload.PowerType,
		PricePerDay:     payload.PricePerDay,
		PricePerHour:    payload.PricePerHour,
		VehicleType:     payload.VehicleType,
		Images:          payload.Images,
		LuggageCapacity: payload.LuggageCapacity,

		CreatedAt: updatedVehicle.CreatedAt,
		UpdatedAt: time.Now(),
	}

	vc.DB.Model(&updatedVehicle).Updates(vehicleToUpdate)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedVehicle})
}

// [...] Get Vehicles Handler
func (vc *VehicleController) GetVehicles(ctx *gin.Context) {
	var vehicles []models.Vehicle

	query := utils.ApplyDynamicPreloading(vc.DB, ctx, vehicleAllowedEntities)

	if err := query.Find(&vehicles).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	vehiclesResponse := utils.MapVehiclesToVehicleResponses(&vehicles)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"vehicles": vehiclesResponse, "count": len(vehiclesResponse)}})
}

// [...] Get Vehicle by ID Handler
func (vc *VehicleController) GetVehicleByID(ctx *gin.Context) {
	var vehicle models.Vehicle
	id := ctx.Param("id")

	query := utils.ApplyDynamicPreloading(vc.DB, ctx, vehicleAllowedEntities)

	if err := query.Where("uuid = ?", id).First(&vehicle).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	vehicleResponse := utils.MapVehicleToVehicleResponse(&vehicle)
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": vehicleResponse})
}

// [...] Delete Post Handler
func (vc *VehicleController) DeleteVehicle(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	vehicleId := ctx.Param("vehicleId")

	var vehicle models.Vehicle
	result := vc.DB.First(&vehicle, "uuid = ?", vehicleId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No vehicle with indicated ID exists"})
		return
	}

	if vehicle.OwnerID != currentUser.ID {
		ctx.JSON(http.StatusForbidden, gin.H{"status": "fail", "message": "You are not authorized to delete this vehicle"})
		return
	}

	vc.DB.Delete(&models.Vehicle{}, "uuid = ?", vehicleId)

	ctx.JSON(http.StatusNoContent, nil)
}
