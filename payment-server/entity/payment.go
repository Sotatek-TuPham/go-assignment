package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	ID      uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	OrderID uuid.UUID `json:"orderId"`
	Status  string    `json:"status"`
}
