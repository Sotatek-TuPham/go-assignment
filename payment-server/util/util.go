package util

import (
	"math/rand"
	"time"
)

// Define PaymentStatus type
type PaymentStatus string

const (
	CANCELLED PaymentStatus = "CANCELLED"
	CONFIRMED PaymentStatus = "CONFIRMED"
)

// Function to get a random payment status
func RandomPaymentStatus() PaymentStatus {
	statuses := []PaymentStatus{CANCELLED, CONFIRMED}
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(statuses))
	return statuses[randomIndex]
}
