package utility

import (
	"math/rand"
	"time"
)

func GeneratePIN() string {
	rand.Seed(time.Now().UnixNano())
	letters := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	pin := make([]rune, 4)
	for i := range pin {
		pin[i] = letters[rand.Intn(len(letters))]
	}
	return string(pin)
}
