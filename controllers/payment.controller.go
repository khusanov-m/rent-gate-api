package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/khusanov-m/rent-gate-api/models"
	"github.com/khusanov-m/rent-gate-api/utils"
	"gorm.io/gorm"
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

	var payments []models.Payment
	results := query.Limit(pagination.Limit).Offset(pagination.Offset).Find(&payments)
	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	//postsResponse := utils.MapPaymentsToPayments(&payments)
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(payments), "data": payments})
}

func (pc *PaymentController) CreatePayment(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)

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
		amount = vehicle.PricePerHour / 24 * hours
	}

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

	// TODO: Add to the history of payments

	if result.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
	}
}
