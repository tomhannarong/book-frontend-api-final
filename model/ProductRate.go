package model

import (
	"time"

	"gorm.io/gorm"
)

// ProductRate - Model
type ProductRate struct {
	gorm.Model
	ID        uint      `json:"id" gorm:"primary_key"`
	ProductID int64     `json:"product_id"`
	Product   Product   `gorm:"-;references:ProductID"`
	Username  string    `json:"username"`
	User      User      `gorm:"-;foreignKey:Username;references:Username"`
	Rate      float64   `json:"rate"`
	CreatedAt time.Time `json:"time"`
}
