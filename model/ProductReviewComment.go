package model

import (
	"time"

	"gorm.io/gorm"
)

// ProductReviewComment - Model
type ProductReviewComment struct {
	gorm.Model
	ID                     uint          `json:"id" gorm:"primary_key"`
	ProductID              int64         `json:"product_id"`
	Product                Product       `gorm:"-;references:ProductID"`
	ProductReviewCommentID int64         `json:"product_review_comment_id"`
	ProductReview          ProductReview `gorm:"-;references:ProductReviewCommentID"`
	Username               string        `json:"username"`
	User                   User          `gorm:"-;foreignKey:Username;references:Username"`
	Message                string        `json:"message"`
	CreatedAt              time.Time     `json:"time"`
}
