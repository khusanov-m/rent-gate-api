package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/khusanov-m/rent-gate-api/models"
	"github.com/khusanov-m/rent-gate-api/utils"
	"gorm.io/gorm"
	"math"
	"net/http"
)

var historyAllowedEntities utils.PreloadEntities = utils.PreloadEntities{
	"PaymentDetails": true,
}

type HistoryController struct {
	DB *gorm.DB
}

func NewHistoryController(DB *gorm.DB) HistoryController {
	return HistoryController{DB}
}

//func (hc *HistoryController) GetAllHistoryRecords(ctx *gin.Context) {
//	pagination, err := utils.NewPaginationFromQuery(ctx)
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"status": "", "message": err.Error()})
//		return
//	}
//
//	query := utils.ApplyDynamicPreloading(hc.DB, ctx, historyAllowedEntities) // payments?preload=PaymentDetails
//
//	var totalItems int64
//	var rentals []models.RentPaymentHistory
//	if err := query.Model(&models.RentPaymentHistory{}).Count(&totalItems).Error; err != nil {
//		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to count total items"})
//		return
//	}
//
//	results := query.Limit(pagination.Limit).Offset(pagination.Offset).Find(&rentals)
//	if results.Error != nil {
//		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error.Error()})
//		return
//	}
//
//	rentalRecords := utils.MapRentalsHistoryToRentalsHistoryResponse(&rentals)
//	totalPages := int(math.Ceil(float64(totalItems) / float64(pagination.Limit)))
//
//	ctx.JSON(http.StatusOK, gin.H{
//		"status": "success",
//		"data": gin.H{
//			"rental_history": rentalRecords,
//			"pagination": gin.H{
//				"total_pages": totalPages,
//				"total_items": totalItems,
//			},
//		},
//	})
//}

//func (hc *HistoryController) GetAllSubscriptionRecords(ctx *gin.Context) {}

func (hc *HistoryController) GetAllRentalRecords(ctx *gin.Context) {
	pagination, err := utils.NewPaginationFromQuery(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "", "message": err.Error()})
		return
	}

	query := utils.ApplyDynamicPreloading(hc.DB, ctx, historyAllowedEntities) // payments?preload=PaymentDetails

	var totalItems int64
	var rentals []models.RentPaymentHistory
	if err := query.Model(&models.RentPaymentHistory{}).Count(&totalItems).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to count total items"})
		return
	}

	results := query.Limit(pagination.Limit).Offset(pagination.Offset).Find(&rentals)

	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error.Error()})
		return
	}

	// filter &rentals by user_id
	currentUser := ctx.MustGet("currentUser").(models.User)

	var filteredRentals []models.RentPaymentHistory
	for _, rental := range rentals {
		if rental.UserID == currentUser.ID {
			filteredRentals = append(filteredRentals, rental)
		}

	}

	rentalRecords := utils.MapRentalsHistoryToRentalsHistoryResponse(&filteredRentals)
	totalPages := int(math.Ceil(float64(totalItems) / float64(pagination.Limit)))

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": gin.H{
			"rental_history": rentalRecords,
			"pagination": gin.H{
				"total_pages": totalPages,
				"total_items": totalItems,
			},
		},
	})
}

func (hc *HistoryController) GetRentalRecordByID(ctx *gin.Context) {
	var rentalRecord models.RentPaymentHistory
	id := ctx.Param("rentId")
	query := utils.ApplyDynamicPreloading(hc.DB, ctx, historyAllowedEntities)

	if err := query.Where("uuid = ?", id).First(&rentalRecord).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	currentUser := ctx.MustGet("currentUser").(models.User)
	if rentalRecord.UserID != currentUser.ID {
		ctx.JSON(http.StatusForbidden, gin.H{"status": "error", "message": "You are not allowed to view this record"})
		return
	}
	rentalResponse := utils.MapRentalHistoryToRentalHistoryResponse(&rentalRecord)
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": rentalResponse})
}

//func (hc *HistoryController) GetSubscriptionByID(ctx *gin.Context) {}

// delete method
func (hc *HistoryController) DeleteRentalRecord(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)

	rentId := ctx.Param("rentId")
	var rentalRecord models.RentPaymentHistory

	result := hc.DB.First(&rentalRecord, "uuid = ?", rentId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No rental record with indicated ID exists"})
		return
	}

	if rentalRecord.UserID != currentUser.ID {
		ctx.JSON(http.StatusForbidden, gin.H{"status": "fail", "message": "You are not authorized to delete this rental record"})
		return
	}

	hc.DB.Delete(&models.RentPaymentHistory{}, "uuid = ?", rentId)

	ctx.JSON(http.StatusNoContent, nil)
}
