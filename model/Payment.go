package model

import (
	"time"

	"gorm.io/gorm"
)

// Payment - Model
type Payment struct {
	gorm.Model
	ID                 uint      `json:"id" gorm:"primary_key"`
	Payment            string    `json:"payment"`
	PaymentDescription string    `json:"payment_description"`
	CreatedAt          time.Time `json:"time"`
}
