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

	HistoryController      controllers.HistoryController
	HistoryRouteController routes.HistoryRouteController
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

	HistoryController = controllers.NewHistoryController(initializers.DB)
	HistoryRouteController = routes.NewHistoryRouteController(HistoryController)

	if config.Environment == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	server = gin.Default()
	server.Use(corsMiddleware())
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8000", "http://localhost:3000", config.ClientOrigin}
	corsConfig.AllowCredentials = true
	server.Use(cors.New(corsConfig))
	/*
		Remove the below block if the production won't work
	*/
	// Disabling all proxy access for security reasons
	server.ForwardedByClientIP = true
	err = server.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		return
	}
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}
	router := server.Group("/api/v1")
	// PING method to check service status
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "pong"})
	})

	AuthRouteController.AuthRoute(router)
	UserRouteController.UserRoute(router)
	VehiclesRouteController.VehicleRoute(router)
	PaymentRouteController.PaymentRoute(router)
	HistoryRouteController.HistoryRoute(router)
	log.Fatal(server.Run(":" + config.ServerPort))
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
