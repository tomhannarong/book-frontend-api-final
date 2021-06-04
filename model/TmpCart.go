package model

import (
	"time"

	"gorm.io/gorm"
)

// TmpCart - Model
type TmpCart struct {
	gorm.Model
	ID             uint      `json:"id" gorm:"primary_key"`
	BookID         int64     `json:"book_id"`
	Product        Product   `gorm:"-;references:BookID"`
	Username       string    `json:"username"`
	UserID         int64     `json:"user_id"`
	User           User      `gorm:"-;foreignKey:Username;references:Username"`
	Quantity       int64     `json:"quantity"`
	IsblameProduct string    `json:"blame_product"`
	IsBuffet       bool      `json:"buffet"`
	IsDiscount     bool      `json:"can_discount"`
	IsActive       bool      `json:"is_active"`
	IsEbook        bool      `json:"is_ebook"`
	CreatedAt      time.Time `json:"time"`
}
