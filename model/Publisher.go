package model

import (
	"time"

	"gorm.io/gorm"
)

// Publisher - Model
type Publisher struct {
	gorm.Model
	ID        uint      `json:"id" gorm:"primary_key"`
	Publisher string    `json:"publisher"`
	CreatedAt time.Time `json:"time"`
}
