package model

import (
	"time"

	"gorm.io/gorm"
)

// Amphure - Model
type Amphure struct {
	gorm.Model
	ID         uint      `json:"id" gorm:"primary_key"`
	Zipcode    string    `json:"zip_code"`
	NameTH     string    `json:"name_th"`
	NameEN     string    `json:"name_en"`
	ProvinceID int64     `json:"ProvinceID"`
	CreatedAt  time.Time `json:"time"`
}
