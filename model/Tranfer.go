package model

import (
	"time"

	"gorm.io/gorm"
)

// Tranfer - Model
type Tranfer struct {
	gorm.Model
	ID             uint      `json:"id" gorm:"primary_key"`
	Amount         float64   `json:"amount"`
	AccountTranfer string    `json:"account_tranfer"`
	Order          OrderMas  `gorm:"-;references:AccountTranfer"`
	BankTranfer    string    `json:"bank_tranfer"`
	Username       string    `json:"username"`
	User           User      `gorm:"-;foreignKey:Username;references:Username"`
	Remark1        string    `json:"remark1"`
	Reason         string    `json:"reason"`
	TranferStatus  string    `json:"tranfer_status"`
	ShowStatus     string    `json:"show_status"`
	TranferDate    time.Time `json:"tranfer_date"`
	CreatedAt      time.Time `json:"time"`
}
