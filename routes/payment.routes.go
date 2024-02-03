package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/khusanov-m/rent-gate-api/controllers"
	"github.com/khusanov-m/rent-gate-api/middleware"
)

type PaymentRouteController struct {
	paymentController controllers.PaymentController
}

func NewPaymentRouteController(paymentController controllers.PaymentController) PaymentRouteController {
	return PaymentRouteController{paymentController}
}

func (pc *PaymentRouteController) PaymentRoute(rg *gin.RouterGroup) {
	router := rg.Group("payments")
	router.Use(middleware.DeserializeUser())
	router.GET("/", pc.paymentController.GetAllPayments)
	router.GET("/:paymentId", pc.paymentController.GetPaymentByID)
	router.POST("/:vehicleId", pc.paymentController.CreatePayment)
	router.POST("/:vehicleId/confirm", pc.paymentController.ConfirmPayment)
	router.PUT("/:paymentId", pc.paymentController.CancelPayment)
	router.DELETE("/:paymentId", pc.paymentController.DeletePayment)
}
