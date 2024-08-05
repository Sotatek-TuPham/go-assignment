package dto

type CreateOrder struct {
	Email    string `json:"email"`
	Quantity uint   `json:"quantity"`
}
