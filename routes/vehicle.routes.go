package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/khusanov-m/rent-gate-api/controllers"
	"github.com/khusanov-m/rent-gate-api/middleware"
)

type VehicleRouteController struct {
	vehicleController controllers.VehicleController
}

func NewRouteVehicleController(vehicleController controllers.VehicleController) VehicleRouteController {
	return VehicleRouteController{vehicleController}
}

func (vc *VehicleRouteController) VehicleRoute(rg *gin.RouterGroup) {
	router := rg.Group("/vehicles")
	router.Use(middleware.DeserializeUser())
	router.GET("", vc.vehicleController.GetVehicles)
	router.GET("/:id", vc.vehicleController.GetVehicleByID)
	router.POST("", vc.vehicleController.CreateVehicle)
	router.PUT("/:vehicleId", vc.vehicleController.UpdateVehicle)
	router.DELETE("/:vehicleId", vc.vehicleController.DeleteVehicle)

}
