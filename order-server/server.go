package main

import (
	"gin-server/config"
	"gin-server/controller"
	"gin-server/service"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	orderService    service.OrderService       = service.NewOrderService()
	orderController controller.OrderController = controller.NewOrderController(orderService)
)

func main() {
	server := gin.Default()
	godotenv.Load()
	config.Connect()
	server.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "first time using gin",
		})
	})

	server.POST("/orders", func(ctx *gin.Context) {
		ctx.JSON(200, orderController.CreateOrder(ctx))
	})

	server.POST("/orders/payment-hook", func(ctx *gin.Context) {
		ctx.JSON(200, orderController.PaymentHook(ctx))
	})

	server.PUT("/orders/:id/cancel", orderController.CancelOrder)

	server.Run(os.Getenv("PORT"))
}
