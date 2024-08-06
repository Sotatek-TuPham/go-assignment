package service

import (
	"gin-server/dto"
	"gin-server/entity"
)

type PaymentService interface {
	CreateOrder(dto.CreatePayment) entity.Payment
}

type paymentService struct{}

func NewPaymentService() PaymentService {
	return &paymentService{}
}

func (service *paymentService) CreateOrder(orderParams dto.CreatePayment) entity.Payment {
	var payment entity.Payment

	return payment
}
