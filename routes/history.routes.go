package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/khusanov-m/rent-gate-api/controllers"
	"github.com/khusanov-m/rent-gate-api/middleware"
)

type HistoryRouteController struct {
	historyController controllers.HistoryController
}

func NewHistoryRouteController(historyController controllers.HistoryController) HistoryRouteController {
	return HistoryRouteController{historyController}
}

func (hc *HistoryRouteController) HistoryRoute(rg *gin.RouterGroup) {
	router := rg.Group("history")
	router.Use(middleware.DeserializeUser())
	//router.GET("/", hc.historyController.GetAllHistoryRecords)
	//
	//subs := router.Group("subscription")
	//subs.GET("/", hc.historyController.GetAllSubscriptionRecords)
	//subs.GET("/:subscriptionId", hc.historyController.GetSubscriptionByID)

	rent := router.Group("rental")
	rent.GET("", hc.historyController.GetAllRentalRecords)
	rent.GET("/:rentId", hc.historyController.GetRentalRecordByID)
	rent.DELETE("/:rentId", hc.historyController.DeleteRentalRecord)
}
