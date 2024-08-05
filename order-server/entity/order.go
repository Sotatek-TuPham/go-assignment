package entity

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Status type
type Status string

const (
	CREATED   Status = "CREATED"
	CONFIRMED Status = "CONFIRMED"
	CANCELLED Status = "CANCELLED"
	DELIVERED Status = "DELIVERED"
)

// Validate the status
func (s Status) IsValid() error {
	switch s {
	case CREATED, CONFIRMED, CANCELLED, DELIVERED:
		return nil
	}
	return errors.New("invalid status")
}

type Order struct {
	gorm.Model
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Email      string    `json:"email"`
	Quantity   uint      `json:"quantity"`
	UnitPrice  uint      `json:"unitPrice"`
	TotalPrice int64     `json:"totalPrice"`
	Status     Status    `gorm:"type:status" json:"status"`
}

// Hook into GORM callbacks
func (u *Order) BeforeSave(tx *gorm.DB) (err error) {
	if err := u.Status.IsValid(); err != nil {
		return err
	}
	return nil
}
