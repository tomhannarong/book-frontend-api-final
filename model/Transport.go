package model

import (
	"time"

	"gorm.io/gorm"
)

// Transport - Model
type Transport struct {
	gorm.Model
	ID                   uint      `json:"id" gorm:"primary_key"`
	TransportRate        float64   `json:"transport_rate"`
	Transport            string    `json:"transport"`
	TransportDescription string    `json:"transport_description"`
	CreatedAt            time.Time `json:"time"`
}
