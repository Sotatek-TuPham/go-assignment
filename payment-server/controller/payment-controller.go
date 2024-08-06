package controller

import (
	"gin-server/dto"
	"gin-server/entity"
	"gin-server/service"

	"github.com/gin-gonic/gin"
)

type PaymentController interface {
	CreateOrder(ctx *gin.Context) entity.Payment
}

type paymentController struct {
	service service.PaymentService
}

func NewPaymentController(service service.PaymentService) PaymentController {
	return &paymentController{
		service: service,
	}
}

func (controller *paymentController) CreateOrder(ctx *gin.Context) entity.Payment {
	var orderParams dto.CreatePayment
	ctx.BindJSON(&orderParams)
	order := controller.service.CreateOrder(orderParams)
	return order
}
