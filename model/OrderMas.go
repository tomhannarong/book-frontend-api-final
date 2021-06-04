package model

import (
	"time"

	"gorm.io/gorm"
)

// OrderMas - Model
type OrderMas struct {
	gorm.Model
	ID             uint      `json:"id" gorm:"primary_key"`
	Username       string    `json:"username"`
	Payment        string    `json:"payment"`
	Transport      string    `json:"transport"`
	TransportRate  float64   `json:"transport_rate"`
	NetPrice       float64   `json:"net_price"`
	OrderStatus    string    `json:"order_status"`
	Reason         string    `json:"reason"`
	BookAvailable  string    `json:"book_available"`
	ApproveStatus  string    `json:"approve_status"`
	ShowStatus     string    `json:"show_status"`
	TranferAddress string    `json:"tranfer_address"`
	TrackingNumber string    `json:"tracking_number"`
	Confirmation   string    `json:"confirmation"`
	AddressID      int64     `json:"address_id"`
	UserID         int64     `json:"user_id"`
	User           User      `gorm:"-;references:UserID"`
	Address        []Address `gorm:"-;references:AddressID"`
	ApproveDate    time.Time `json:"approve_date"`
	DatePaid       time.Time `json:"date_paid"`
	CreatedAt      time.Time `json:"time"`
}
