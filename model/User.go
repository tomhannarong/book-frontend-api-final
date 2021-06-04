package model

import (
	"time"

	"gorm.io/gorm"
)

// User - Model
type User struct {
	gorm.Model
	ID        uint   `json:"id" gorm:"primary_key"`
	Username  string `gorm:"unique" form:"username" binding:"required"`
	Password  string `form:"password" binding:"required"`
	Level     string `gorm:"default:normal"`
	CreatedAt time.Time
}

// https://gorm.io/docs/models.html
