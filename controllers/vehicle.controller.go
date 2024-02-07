package controllers

// CAN PRELOAD: User

import (
	"math"
	"net/http"
	"strconv"
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
	pagination, err := utils.NewPaginationFromQuery(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	// Get filtering parameters from query string
	priceFromStr := ctx.Query("price_from")
	priceToStr := ctx.Query("price_to")
	vehicleType := ctx.Query("vehicle_type")

	// Convert price range parameters to float64
	var priceFrom, priceTo float64
	if priceFromStr != "" {
		priceFrom, err = strconv.ParseFloat(priceFromStr, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid price_from parameter"})
			return
		}
	}
	if priceToStr != "" {
		priceTo, err = strconv.ParseFloat(priceToStr, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid price_to parameter"})
			return
		}
	}

	query := vc.DB.Model(&models.Vehicle{})

	query = utils.ApplyDynamicPreloading(vc.DB, ctx, vehicleAllowedEntities)

	// Apply filters
	if priceFrom > 0 && priceTo > 0 && priceFrom <= priceTo {
		query = query.Where("price_per_day BETWEEN ? AND ?", priceFrom, priceTo)
	} else {
		if priceFrom > 0 {
			query = query.Where("price_per_day >= ?", priceFrom)
		}

		if priceTo > 0 && priceTo >= priceFrom {
			query = query.Where("price_per_day <= ?", priceTo)
		}
	}
	if vehicleType != "" {
		query = query.Where("vehicle_type = ?", vehicleType)
	}

	var totalItems int64
	var vehicles []models.Vehicle

	if err := query.Model(&models.Vehicle{}).Count(&totalItems).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to count total items"})
		return
	}

	results := query.Limit(pagination.Limit).Offset(pagination.Offset).Find(&vehicles)
	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error.Error()})
		return
	}

	vehiclesResponse := utils.MapVehiclesToVehicleResponses(&vehicles)

	totalPages := int(math.Ceil(float64(totalItems) / float64(pagination.Limit)))

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": gin.H{
			"vehicles": vehiclesResponse,
			"count":    len(vehiclesResponse),
			"pagination": gin.H{
				"total_pages": totalPages,
				"total_items": totalItems,
			},
		},
	})
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
