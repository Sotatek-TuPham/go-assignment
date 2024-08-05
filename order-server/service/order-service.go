package service

import (
	"gin-server/config"
	"gin-server/dto"
	"gin-server/entity"
	"math/rand"
)

type OrderService interface {
	CreateOrder(dto.CreateOrder) entity.Order
}

type orderService struct{}

func NewOrderService() OrderService {
	return &orderService{}
}

func (service *orderService) CreateOrder(orderParams dto.CreateOrder) entity.Order {
	var order entity.Order
	order.Email = orderParams.Email
	order.Quantity = orderParams.Quantity
	order.UnitPrice = uint(rand.Intn(100))
	order.TotalPrice = int64(order.UnitPrice) * int64(order.Quantity)
	order.Status = entity.CREATED
	config.DB.Create(&order)

	return order
}
