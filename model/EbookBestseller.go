package model

import (
	"time"

	"gorm.io/gorm"
)

// EbookBestseller - Model
type EbookBestseller struct {
	gorm.Model
	ID        uint      `json:"id" gorm:"primary_key"`
	ProductID int64     `json:"product_id"`
	Product   Product   `gorm:"-;references:ProductID"`
	Seq       int64     `json:"seq"`
	PostDate  time.Time `json:"post_date"`
	CreatedAt time.Time `json:"time"`
}
