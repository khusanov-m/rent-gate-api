package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/khusanov-m/rent-gate-api/controllers"
	"github.com/khusanov-m/rent-gate-api/initializers"
	"github.com/khusanov-m/rent-gate-api/routes"
)

var (
	server              *gin.Engine
	AuthController      controllers.AuthController
	AuthRouteController routes.AuthRouteController

	UserController      controllers.UserController
	UserRouteController routes.UserRouteController

	VehiclesController      controllers.VehicleController
	VehiclesRouteController routes.VehicleRouteController

	PaymentController      controllers.PaymentController
	PaymentRouteController routes.PaymentRouteController
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)

	AuthController = controllers.NewAuthController(initializers.DB)
	AuthRouteController = routes.NewAuthRouteController(AuthController)

	UserController = controllers.NewUserController(initializers.DB)
	UserRouteController = routes.NewRouteUserController(UserController)

	VehiclesController = controllers.NewVehicleController(initializers.DB)
	VehiclesRouteController = routes.NewRouteVehicleController(VehiclesController)

	PaymentController = controllers.NewPaymentController(initializers.DB)
	PaymentRouteController = routes.NewPaymentRouteController(PaymentController)

	if config.Environment == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	server = gin.Default()
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8000", "http://localhost:3000", config.ClientOrigin}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	router := server.Group("/api/v1")
	// PING method to check service status
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "pong"})
	})

	AuthRouteController.AuthRoute(router)
	UserRouteController.UserRoute(router)
	VehiclesRouteController.VehicleRoute(router)
	PaymentRouteController.PaymentRoute(router)
	log.Fatal(server.Run(":" + config.ServerPort))
}
