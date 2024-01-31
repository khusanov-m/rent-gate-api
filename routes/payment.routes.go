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
	router.GET("/", middleware.DeserializeUser(), pc.paymentController.GetAllPayments)
	//router.GET("/", middleware.DeserializeUser(), pc.paymentController.GetPaymentByID)
	//router.POST("/", middleware.DeserializeUser(), pc.paymentController.CreatePayment)
	//router.PUT("/:paymentId", middleware.DeserializeUser(), pc.paymentController.UpdatePayment)
	//router.DELETE("/", middleware.DeserializeUser(), pc.paymentController.DeletePayment)
}
