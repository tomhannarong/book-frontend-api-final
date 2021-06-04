package model

import (
	"time"

	"gorm.io/gorm"
)

// BoardPost - Model
type BoardPost struct {
	gorm.Model
	ID              uint      `json:"id" gorm:"primary_key"`
	Topic           string    `json:"topic"`
	Username        string    `json:"username"`
	User            User      `gorm:"-;references:Username"`
	View            int64     `json:"view"`
	PostDescription string    `json:"post_description"`
	ShowStatus      string    `json:"show_status"`
	Pin             string    `json:"pin"`
	CreatedAt       time.Time `json:"time"`
}
