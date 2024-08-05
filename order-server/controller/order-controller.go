package controller

import (
	"gin-server/dto"
	"gin-server/entity"
	"gin-server/service"

	"github.com/gin-gonic/gin"
)

type OrderController interface {
	CreateOrder(ctx *gin.Context) entity.Order
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
