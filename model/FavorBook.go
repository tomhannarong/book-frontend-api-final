package model

import (
	"time"

	"gorm.io/gorm"
)

// FavorBook - Model
type FavorBook struct {
	gorm.Model
	ID        uint      `json:"id" gorm:"primary_key"`
	BookID    int64     `json:"book_id"`
	Product   Product   `gorm:"-;references:BookID"`
	UserID    int64     `json:"user_id"`
	User      User      `gorm:"-;references:UserID"`
	IsEbook   string    `json:"is_ebook"`
	CreatedAt time.Time `json:"time"`
}
