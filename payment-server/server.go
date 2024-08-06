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
	orderService    service.PaymentService       = service.NewPaymentService()
	orderController controller.PaymentController = controller.NewPaymentController(orderService)
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

	server.POST("/payments", func(ctx *gin.Context) {
		ctx.JSON(200, orderController.CreatePayment(ctx))
	})

	server.Run(os.Getenv("PORT"))
}
