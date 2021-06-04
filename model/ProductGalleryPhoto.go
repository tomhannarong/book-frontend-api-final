package model

import (
	"time"

	"gorm.io/gorm"
)

// ProductGalleryPhoto - Model
type ProductGalleryPhoto struct {
	gorm.Model
	ID            uint      `json:"id" gorm:"primary_key"`
	ProductID     int64     `json:"product_id"`
	Product       Product   `gorm:"-;references:ProductID"`
	Photo         string    `json:"photo"`
	DefaultActive bool      `json:"default default:0"`
	CreatedAt     time.Time `json:"time"`
}
