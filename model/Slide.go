package model

import (
	"time"

	"gorm.io/gorm"
)

// Slide - Model
type Slide struct {
	gorm.Model
	ID          uint      `json:"id" gorm:"primary_key"`
	Position    int64     `json:"position"`
	SlideName   string    `json:"slide_name"`
	SlideImages string    `json:"slide_images"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"time"`
}
