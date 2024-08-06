package controller

import (
	"gin-server/dto"
	"gin-server/service"

	"github.com/gin-gonic/gin"
)

type PaymentController interface {
	CreatePayment(ctx *gin.Context) dto.PaymentResponse
}

type paymentController struct {
	service service.PaymentService
}

func NewPaymentController(service service.PaymentService) PaymentController {
	return &paymentController{
		service: service,
	}
}

func (controller *paymentController) CreatePayment(ctx *gin.Context) dto.PaymentResponse {
	var orderParams dto.CreatePayment
	ctx.BindJSON(&orderParams)
	response := controller.service.CreatePayment(orderParams)
	return response
}
