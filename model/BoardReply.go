package model

import (
	"time"

	"gorm.io/gorm"
)

// BoardReply - Model
type BoardReply struct {
	gorm.Model
	ID        uint      `json:"id" gorm:"primary_key"`
	TopicID   int64     `json:"topic_id"`
	Username  string    `json:"username"`
	User      User      `gorm:"-;references:Username"`
	IPUser    string    `json:"ip_user"`
	Reply     string    `json:"reply"`
	CreatedAt time.Time `json:"time"`
}
