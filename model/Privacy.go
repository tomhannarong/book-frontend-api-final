package model

import (
	"time"

	"gorm.io/gorm"
)

// Privacy - Model
type Privacy struct {
	gorm.Model
	ID            uint      `json:"id" gorm:"primary_key"`
	Title         string    `json:"title"`
	Detail        string    `json:"detail"`
	ApproveStatus bool      `json:"approve_status"`
	CreatedAt     time.Time `json:"time"`
}
