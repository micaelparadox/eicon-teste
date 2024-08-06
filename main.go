package main

import (
	"test-eicon/config"
	"test-eicon/controllers"
	"test-eicon/repositories"
	"test-eicon/services"
	"test-eicon/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			config.LoadConfig,
			utils.NewDatabase,
			repositories.NewOrderRepository,
			services.NewOrderService,
			controllers.NewOrderController,
			gin.Default,
		),
		fx.Invoke(
			registerRoutes,
		),
	)
	app.Run()
}

func registerRoutes(router *gin.Engine, orderController *controllers.OrderController) {
	router.POST("/orders", orderController.CreateOrder)
	router.GET("/orders", orderController.GetOrders)
	router.Run(":8080")
}
