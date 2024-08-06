package service

import (
	"gin-server/config"
	"gin-server/dto"
	"gin-server/entity"
)

type PaymentService interface {
	CreatePayment(dto.CreatePayment) dto.PaymentResponse
}

type paymentService struct{}

func NewPaymentService() PaymentService {
	return &paymentService{}
}

func (service *paymentService) CreatePayment(orderParams dto.CreatePayment) dto.PaymentResponse {
	var payment entity.Payment
	payment.OrderID = orderParams.OrderId
	payment.Status = "CREATED"
	config.DB.Create(&payment)
	return dto.PaymentResponse{
		Message: "Payment Created",
	}
}
