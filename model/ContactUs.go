package model

import (
	"time"

	"gorm.io/gorm"
)

// ContactUs - Model
type ContactUs struct {
	gorm.Model
	ID        uint      `json:"id" gorm:"primary_key"`
	Topic     string    `json:"topic"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Tel       string    `json:"tel"`
	Message   string    `json:"message"`
	Read      string    `json:"read default:'false'"`
	CreatedAt time.Time `json:"time"`
}
