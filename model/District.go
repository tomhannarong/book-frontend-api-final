package model

import (
	"time"

	"gorm.io/gorm"
)

// District - Model
type District struct {
	gorm.Model
	ID        uint      `json:"id" gorm:"primary_key"`
	Zipcode   string    `json:"zip_code"`
	NameTH    string    `json:"name_th"`
	NameEN    string    `json:"name_en"`
	AmphureID int64     `json:"amphure_id"`
	CreatedAt time.Time `json:"time"`
}
