package dto

type CreatePayment struct {
	OrderId string `json:"orderId"`
	PIN     string `json:"pin"`
}

type PaymentResponse struct {
	Message string `json:"message"`
}
