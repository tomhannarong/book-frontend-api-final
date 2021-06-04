package model

import (
	"time"

	"gorm.io/gorm"
)

// OrderHistory - Model
type OrderHistory struct {
	gorm.Model
	ID              uint       `json:"id" gorm:"primary_key"`
	ProductID       int64      `json:"product_id"`
	OrderID         int64      `json:"order_id"`
	User            User       `gorm:"-;foreignKey:Username;references:Username"`
	Products        []*Product `gorm:"-;references:ProductID"`
	Order           OrderMas   `gorm:"-;references:OrderID"`
	Username        string     `json:"username"`
	Price           float64    `json:"price"`
	ProductPrice    float64    `json:"product_price"`
	Quantitys       int64      `json:"quantitys"`
	PercentDiscount int64      `json:"percent_discount"`
	Discount        float64    `json:"discount"`
	Net             float64    `json:"net"`
	SharePercent    int64      `json:"share_percent"`
	Buffet          string     `json:"buffet"`
	IsEbook         string     `json:"is_ebook"`
	CreatedAt       time.Time  `json:"time"`
}
