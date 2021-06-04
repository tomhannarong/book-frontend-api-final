package model

import (
	"time"

	"gorm.io/gorm"
)

// EbookApproveEbook - Model
type EbookApproveEbook struct {
	gorm.Model
	ID            uint      `json:"id" gorm:"primary_key"`
	Username      string    `json:"username"`
	ProductID     int64     `json:"product_id"`
	Product       Product   `gorm:"-;references:ProductID"`
	OrderID       int64     `json:"order_id"`
	OrderMas      OrderMas  `gorm:"-;references:OrderID"`
	TranID        int64     `json:"tran_id"`
	OrderTran     OrderTran `gorm:"-;references:TranID"`
	ApproveStatus string    `json:"approve_status"`
	ApproveDate   time.Time `json:"ApproveDate"`
	CreatedAt     time.Time `json:"time"`
}
