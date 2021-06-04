package model

import (
	"time"

	"gorm.io/gorm"
)

// Province - Model
type Province struct {
	gorm.Model
	ID        uint      `json:"id" gorm:"primary_key"`
	NameTH    string    `json:"name_th"`
	NameEN    string    `json:"name_en"`
	CreatedAt time.Time `json:"time"`
}
