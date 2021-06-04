package model

import (
	"time"

	"gorm.io/gorm"
)

// ProductReview - Model
type ProductReview struct {
	gorm.Model
	ID                    uint                    `json:"id" gorm:"primary_key"`
	ProductID             int64                   `json:"product_id"`
	Product               Product                 `gorm:"-;references:ProductID"`
	Username              string                  `json:"username"`
	User                  User                    `gorm:"-;foreignKey:Username;references:Username"`
	Message               string                  `json:"message"`
	ProductReviewComments []*ProductReviewComment `gorm:"-;foreignKey:ProductReviewCommentID"`
	CreatedAt             time.Time               `json:"time"`
}
