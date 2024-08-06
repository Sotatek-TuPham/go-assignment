package dto

import "github.com/google/uuid"

type CreateOrder struct {
	Email    string `json:"email"`
	Quantity uint   `json:"quantity"`
}

type OrderPayload struct {
	OrderID uuid.UUID `json:"orderId"`
	PIN     string    `json:"pin"`
}
