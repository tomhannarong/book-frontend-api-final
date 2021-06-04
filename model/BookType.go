package model

import (
	"time"

	"gorm.io/gorm"
)

// BookType - Model
type BookType struct {
	gorm.Model
	ID        uint       `json:"id" gorm:"primary_key"`
	BookType  string     `json:"book_type"`
	Products  []*Product `gorm:"-;references:BookTypeID"`
	CreatedAt time.Time  `json:"time"`
}
