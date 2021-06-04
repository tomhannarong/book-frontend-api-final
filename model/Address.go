package model

import (
	"time"

	"gorm.io/gorm"
)

// Address - Model
type Address struct {
	gorm.Model
	ID         uint      `json:"id" gorm:"primary_key"`
	Fname      string    `json:"fname"`
	Lname      string    `json:"lname"`
	Tel        string    `json:"tel"`
	AddressDes string    `json:"addressDes"`
	Subdistric string    `json:"subdistric"`
	Distric    string    `json:"distric"`
	Province   string    `json:"province"`
	Zipcode    string    `json:"zipcode"`
	UserID     string    `json:"UserID"`
	User       User      `gorm:"-;references:UserID"`
	Default    bool      `json:"default default:0"`
	CreatedAt  time.Time `json:"time"`
}
