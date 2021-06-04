package db

import (
	"book-frontend-api-final/config"
	"book-frontend-api-final/model"
	"log"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var db *gorm.DB

// GetDB - call this method to get db
func GetDB() *gorm.DB {
	return db
}

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file ")
	}
}

// SetupDB - setup dabase for sharing to all api
func SetupDB() {

	database := config.DBConfig()

	database.AutoMigrate(&model.User{}) // Not Complete
	database.AutoMigrate(&model.Product{})
	database.AutoMigrate(&model.Transaction{})
	database.AutoMigrate(&model.Address{})
	database.AutoMigrate(&model.BestSeller{})
	database.AutoMigrate(&model.Amphure{})
	database.AutoMigrate(&model.BoardPost{})
	database.AutoMigrate(&model.BoardReply{})
	database.AutoMigrate(&model.BookType{})
	database.AutoMigrate(&model.ContactUs{})
	database.AutoMigrate(&model.District{})
	database.AutoMigrate(&model.EbookApproveEbook{})
	database.AutoMigrate(&model.EbookBestseller{})
	database.AutoMigrate(&model.FavorBook{})
	database.AutoMigrate(&model.ImageSlip{})
	database.AutoMigrate(&model.OrderHistory{})
	database.AutoMigrate(&model.OrderMas{})
	database.AutoMigrate(&model.OrderTran{})
	database.AutoMigrate(&model.Payment{})
	database.AutoMigrate(&model.Privacy{})
	database.AutoMigrate(&model.Province{})
	database.AutoMigrate(&model.Publisher{})
	database.AutoMigrate(&model.Transport{})
	database.AutoMigrate(&model.Slide{})
	database.AutoMigrate(&model.TmpCart{})
	database.AutoMigrate(&model.ProductGalleryPhoto{})
	database.AutoMigrate(&model.ProductRate{})
	database.AutoMigrate(&model.ProductReview{})
	database.AutoMigrate(&model.ProductReviewComment{})
	database.AutoMigrate(&model.Tranfer{})

	db = database
}
