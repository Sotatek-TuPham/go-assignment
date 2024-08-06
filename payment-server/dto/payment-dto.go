package dto

import (
	"gin-server/util"

	"github.com/google/uuid"
)

type CreatePayment struct {
	OrderId uuid.UUID `json:"orderId"`
	PIN     string    `json:"pin"`
}

type PaymentResponse struct {
	Message string `json:"message"`
}

type PaymentPayload struct {
	OrderID uuid.UUID          `json:"orderId"`
	Status  util.PaymentStatus `json:"status"`
}
