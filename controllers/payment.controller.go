package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/khusanov-m/rent-gate-api/models"
	"github.com/khusanov-m/rent-gate-api/utils"
	"gorm.io/gorm"
	"net/http"
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
