package controller

import (
	"gin-server/dto"
	"gin-server/entity"
	"gin-server/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderController interface {
	CreateOrder(ctx *gin.Context) entity.Order
	PaymentHook(ctx *gin.Context) dto.OrderResponse
	CancelOrder(ctx *gin.Context)
}

type orderController struct {
	service service.OrderService
}

func NewOrderController(service service.OrderService) OrderController {
	return &orderController{
		service: service,
	}
}

func (controller *orderController) CreateOrder(ctx *gin.Context) entity.Order {
	var orderParams dto.CreateOrder
	ctx.BindJSON(&orderParams)
	order := controller.service.CreateOrder(orderParams)
	return order
}

func (controller *orderController) PaymentHook(ctx *gin.Context) dto.OrderResponse {
	var paymentPayload dto.PaymentPayload
	ctx.BindJSON(&paymentPayload)
	response := controller.service.PaymentHook(paymentPayload)
	return response
}

func (controller *orderController) CancelOrder(ctx *gin.Context) {
	var orderId string = ctx.Param("id")
	response, err := controller.service.CancelOrder(orderId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response)
}
