package model

import (
	"time"

	"gorm.io/gorm"
)

// ImageSlip - Model
type ImageSlip struct {
	gorm.Model
	ID        uint      `json:"id" gorm:"primary_key"`
	Filename  string    `json:"filename"`
	OrderID   int64     `json:"order_id"`
	Order     OrderMas  `gorm:"-;references:OrderID"`
	CreatedAt time.Time `json:"time"`
}
