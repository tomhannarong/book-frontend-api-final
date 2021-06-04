package model

import (
	"time"

	"gorm.io/gorm"
)

// BestSeller - Model
type BestSeller struct {
	gorm.Model
	ID        uint      `json:"id" gorm:"primary_key"`
	Top       int64     `json:"top"`
	BookID    int64     `json:"book_id"`
	Product   Product   `gorm:"-;foreignKey:BookID"`
	CreatedAt time.Time `json:"time"`
}
