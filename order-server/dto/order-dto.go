package dto

import (
	"gin-server/entity"

	"github.com/google/uuid"
)

type CreateOrder struct {
	Email    string `json:"email"`
	Quantity uint   `json:"quantity"`
}

type OrderPayload struct {
	OrderID uuid.UUID `json:"orderId"`
	PIN     string    `json:"pin"`
}

type PaymentPayload struct {
	OrderID uuid.UUID `json:"orderId"`
	Status  string    `json:"status"`
}

type OrderResponse struct {
	Message string `json:"message"`
}

type UpdateOrderDTO struct {
	Status entity.Status `json:"status"`
}
