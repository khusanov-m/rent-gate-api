package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/khusanov-m/rent-gate-api/models"
	"gorm.io/gorm"
)

type VehicleController struct {
	DB *gorm.DB
}

func NewVehicleController(DB *gorm.DB) VehicleController {
	return VehicleController{DB}
}

func (vc *VehicleController) GetVehicles(ctx *gin.Context) {
	var vehicles []models.Vehicle
	if err := vc.DB.Find(&vehicles).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"vehicles": vehicles}})
}

func (vc *VehicleController) GetVehicleByID(ctx *gin.Context) {
	var vehicle models.Vehicle
	id := ctx.Param("id")

	if err := vc.DB.Where("uuid = ?", id).First(&vehicle).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	vehicleResponse := models.VehicleResponse{
		ID:           vehicle.UUID,
		OwnerID:      vehicle.OwnerID,
		CategoryID:   vehicle.CategoryID,
		LocationID:   vehicle.LocationID,
		Status:       vehicle.Status,
		PricePerHour: vehicle.PricePerHour,
		PricePerDay:  vehicle.PricePerDay,
		CreatedAt:    vehicle.CreatedAt,
		UpdatedAt:    vehicle.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"vehicle": vehicleResponse}})
}
