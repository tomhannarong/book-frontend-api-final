package model

import (
	"time"

	"gorm.io/gorm"
)

// Product - Model
type Product struct {
	gorm.Model
	ID                   uint                   `json:"id" gorm:"primary_key"`
	BookName             string                 `json:"book_name"`
	BookTypeID           int64                  `json:"book_type_id"`
	BookType             BookType               `gorm:"-;references:BookTypeID"`
	ProductGalleryPhotos []*ProductGalleryPhoto `gorm:"-;references:ProductID"`
	Writer               string                 `json:"writer"`
	Alias                string                 `json:"alias"`
	Price                float64                `json:"price"`
	ProductPrice         float64                `json:"product_price"`
	Rate                 float64                `json:"rate"`
	RateNum              int64                  `json:"rate_num"`
	Pages                int64                  `json:"pages"`
	BookDescription      string                 `json:"book_description"`
	Attachment           string                 `json:"attachment"`
	OnMarket             string                 `json:"on_market"`
	ISBN                 string                 `json:"ISBN"`
	PimTime              int64                  `json:"pim_time"`
	PimYear              int64                  `json:"pim_year"`
	PublisherID          int64                  `json:"publisher_id"`
	Publisher            Publisher              `gorm:"-;references:PublisherID"`
	TagDescription       string                 `json:"tag_description"`
	TagKeyword           string                 `json:"tag_keyword"`
	IsBlameProduct       bool                   `json:"blame_product"`
	IsSerieProduct       bool                   `json:"serie_product"`
	BlamePosition        string                 `json:"blame_position"`
	BlameImages          string                 `json:"blame_images"`
	IsShowBlame          bool                   `json:"show_blame"`
	BlogURL              string                 `json:"blog_url"`
	YoutubeURL           string                 `json:"youtube_url"`
	Isbuffet             bool                   `json:"buffet"`
	StockTotal           int64                  `json:"stock_total"`
	StockHold            int64                  `json:"stock_hold"`
	StockRemain          int64                  `json:"stock_remain"`
	StockSold            int64                  `json:"stock_sold"`
	PublicShow           bool                   `json:"public_show"`
	IsCanDiscount        bool                   `json:"can_discount"`
	BookWeight           int64                  `json:"book_weight"`
	PromoteLink          string                 `json:"promote_link"`
	ProductPDF           string                 `json:"product_pdf"`
	AffiliatePrice       float64                `json:"affiliate_price"`
	UserID               int64                  `json:"user_id"`
	User                 User                   `gorm:"-;references:UserID"`
	Username             string                 `json:"username"`
	IsPublish            bool                   `json:"publish1"`
	IsBestSeller         bool                   `json:"best_seller"`
	IsRecommended        bool                   `json:"recommended"`
	IsHotItem            bool                   `json:"hot_item"`
	IsUserBook           bool                   `json:"user_book"`
	IsEbook              bool                   `json:"is_ebook	"`
	CreatedAt            time.Time              `json:"time"`
}
