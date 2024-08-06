package dto

type CreatePayment struct {
	OrderId string `json:"orderId"`
}

type PaymentResponse struct {
	Message string `json:"message"`
}
