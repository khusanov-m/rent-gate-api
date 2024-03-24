package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/khusanov-m/rent-gate-api/models"
	"github.com/khusanov-m/rent-gate-api/utils"
	"gorm.io/gorm"
	"math"
	"net/http"
	"strconv"
	"time"
)

//// CAN PRELOAD:

var paymentsAllowedEntities utils.PreloadEntities = utils.PreloadEntities{
	"PaymentDetails": true,
}

type PaymentController struct {
	DB *gorm.DB
}

var paymentAllowedEntities utils.PreloadEntities = utils.PreloadEntities{}

func NewPaymentController(DB *gorm.DB) PaymentController {
	return PaymentController{DB}
}

func (pc *PaymentController) GetAllPayments(ctx *gin.Context) {
	pagination, err := utils.NewPaginationFromQuery(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	query := utils.ApplyDynamicPreloading(pc.DB, ctx, paymentsAllowedEntities) // payments?preload=PaymentDetails

	var totalItems int64
	var payments []models.Payment

	if err := query.Model(&models.Payment{}).Count(&totalItems).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to count total items"})
		return
	}

	results := query.Limit(pagination.Limit).Offset(pagination.Offset).Find(&payments)
	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error.Error()})
		return
	}

	paymentsResponse := utils.MapPaymentsToPaymentsResponse(&payments)

	totalPages := int(math.Ceil(float64(totalItems) / float64(pagination.Limit)))

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": gin.H{
			"payments": paymentsResponse,
			"count":    len(paymentsResponse),
			"pagination": gin.H{
				"total_pages": totalPages,
				"total_items": totalItems,
			},
		},
	})
}

func (pc *PaymentController) GetPaymentByID(ctx *gin.Context) {
	var payment models.Payment
	id := ctx.Param("paymentId")

	query := utils.ApplyDynamicPreloading(pc.DB, ctx, paymentAllowedEntities)

	if err := query.Where("uuid = ?", id).First(&payment).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	paymentRes := utils.MapPaymentToPaymentResponse(&payment)
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": paymentRes})
}

func (pc *PaymentController) CreatePayment(ctx *gin.Context) {
	var payload *models.PaymentInput
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	id := ctx.Param("vehicleId")
	vehicleId, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid vehicle id"})
		return
	}

	var vehicle models.Vehicle
	vehicleRes := pc.DB.First(&vehicle, "uuid = ?", vehicleId)
	if vehicleRes.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No vehicle with indicated ID"})
		return
	}

	var amount float64
	if payload.TotalHours > 23 {
		hours, _ := strconv.ParseFloat(strconv.Itoa(int(payload.TotalHours)), 64)
		amount = vehicle.PricePerDay / 24 * hours
	} else {
		hours, _ := strconv.ParseFloat(strconv.Itoa(int(payload.TotalHours)), 64)
		amount = vehicle.PricePerHour * hours
	}

	amount = amount + payload.AddonsWithDiscountPrice

	currentUser := ctx.MustGet("currentUser").(models.User)
	newPayment := models.Payment{
		UserID:         currentUser.ID,
		Amount:         amount,
		PaymentStatus:  "init",
		PaymentType:    payload.PaymentType,
		PaymentFor:     "rent",
		PaymentDetails: vehicleId,

		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	result := pc.DB.Create(&newPayment)
	if result.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	newRentalRecord := models.RentPaymentHistory{
		UserID:       currentUser.ID,
		VehicleID:    vehicleId,
		PaymentID:    newPayment.UUID,
		TotalAmount:  amount,
		PricePerHour: vehicle.PricePerHour,
		PricePerDay:  vehicle.PricePerDay,
		Duration:     payload.TotalHours,
		Status:       "pending",

		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	rentalRecordResult := pc.DB.Create(&newRentalRecord)
	if rentalRecordResult.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": rentalRecordResult.Error.Error()})
		return
	}

	newPaymentResponse := utils.MapPaymentToPaymentResponse(&newPayment)
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newPaymentResponse})
}

func (pc *PaymentController) DeletePayment(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	paymentId := ctx.Param("paymentId")

	var payment models.Payment

	if result := pc.DB.First(&payment, "uuid = ?", paymentId); result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": result.Error.Error()})
		return
	}

	if payment.UserID != currentUser.ID {
		ctx.JSON(http.StatusForbidden, gin.H{"status": "fail", "message": "You are not authorized to delete this payment"})
		return
	}

	pc.DB.Delete(&models.Payment{}, "uuid = ?", paymentId)
	ctx.JSON(http.StatusNoContent, nil)
}

func (pc *PaymentController) CancelPayment(ctx *gin.Context) {
	paymentId := ctx.Param("paymentId")

	var payment models.Payment
	if result := pc.DB.First(&payment, "uuid = ?", paymentId); result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": result.Error.Error()})
		return
	}

	if payment.PaymentStatus == "cancelled" {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "code": "PAYMENT_ALREADY_CANCELLED", "message": "Payment have been already cancelled"})
		return
	}

	currentUser := ctx.MustGet("currentUser").(models.User)
	if payment.UserID != currentUser.ID {
		ctx.JSON(http.StatusForbidden, gin.H{"status": "fail", "message": "You are not authorized to cancel this payment"})
		return
	}

	// Checking where the payment exist or not
	var paymentRecord models.RentPaymentHistory
	if historyItem := pc.DB.First(&paymentRecord, "payment_id = ?", payment.UUID); historyItem.Error != nil {
		ctx.JSON(http.StatusForbidden, gin.H{"status": "fail", "message": "The payment you are cancelling is not existing"})
		return
	}

	paymentRecordToUpdate := models.RentPaymentHistory{
		Status:    "cancelled",
		UpdatedAt: time.Now(),
	}
	paymentToUpdate := models.Payment{
		PaymentStatus: "cancelled",
		UpdatedAt:     time.Now(),
	}
	pc.DB.Model(&payment).Updates(paymentToUpdate)
	pc.DB.Model(&paymentRecord).Updates(paymentRecordToUpdate)
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": payment})
}

func (pc *PaymentController) ConfirmPayment(ctx *gin.Context) {
	paymentId := ctx.Param("paymentId")

	var payment models.Payment
	if result := pc.DB.First(&payment, "uuid = ?", paymentId); result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": result.Error.Error()})
		return
	}

	if payment.PaymentStatus == "cancelled" {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "code": "CAN_NOT_CONFIRM_CANCELLED", "message": "Cancelled payment can not be confirmed"})
		return
	}

	currentUser := ctx.MustGet("currentUser").(models.User)
	if payment.UserID != currentUser.ID {
		ctx.JSON(http.StatusForbidden, gin.H{"status": "fail", "message": "You are not authorized to cancel this payment"})
		return
	}

	var paymentRecord models.RentPaymentHistory
	if recordItem := pc.DB.First(&paymentRecord, "payment_id = ?", payment.UUID); recordItem.Error != nil {
		ctx.JSON(http.StatusForbidden, gin.H{"status": "fail", "message": recordItem.Error.Error()})
		return
	}

	paymentRecordToUpdate := models.RentPaymentHistory{
		Status:    "confirmed",
		UpdatedAt: time.Now(),
	}
	paymentToUpdate := models.Payment{
		PaymentStatus: "confirmed",
		UpdatedAt:     time.Now(),
	}
	pc.DB.Model(&payment).Updates(paymentToUpdate)
	pc.DB.Model(&paymentRecord).Updates(paymentRecordToUpdate)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": payment})
}
