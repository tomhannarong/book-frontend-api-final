package model

import (
	"time"

	"gorm.io/gorm"
)

// OrderTran - Model
type OrderTran struct {
	gorm.Model
	ID              uint       `json:"id" gorm:"primary_key"`
	OrderID         int64      `json:"order_id"`
	Username        string     `json:"username"`
	User            User       `gorm:"-;foreignKey:Username;references:Username"`
	BookID          int64      `json:"book_id"`
	Products        []*Product `gorm:"-;references:BookID"`
	Quantitys       int64      `json:"quantitys"`
	Price           float64    `json:"price"`
	ProductPrice    float64    `json:"product_price"`
	PercentDiscount int64      `json:"percent_discount"`
	Discount        float64    `json:"discount"`
	Net             float64    `json:"net"`
	SharePercent    float64    `json:"share_percent"`
	ApproveStatus   string     `json:"payment"`
	ApproveDate     time.Time  `json:"approve_date"`
	IsEbook         bool       `json:"is_ebook"`
	CreatedAt       time.Time  `json:"time"`
}
