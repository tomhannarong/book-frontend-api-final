package api

import (
	"book-frontend-api-final/db"
	"book-frontend-api-final/interceptor"
	"book-frontend-api-final/model"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// SetupProductRateAPI - Product Rate Star
func SetupProductRateAPI(router *gin.Engine) {
	productRateAPI := router.Group("/api")
	{
		// Get Product Rate by Id
		productRateAPI.GET("/product/:id/product-rate", interceptor.JwtVerify, getProductRateByID)

		// Create Or Add , Product Rate
		productRateAPI.POST("/product/:id/product-rate", interceptor.JwtVerify, createProductRate)

	}
}

func getProductRateByID(c *gin.Context) {
	var productRate model.ProductRate
	db.GetDB().Where("product_id = ?", c.Param("id")).First(&productRate)
	c.JSON(http.StatusOK, productRate)
}

func createProductRate(c *gin.Context) {
	rate, _ := strconv.ParseFloat(c.PostForm("rate"), 2)
	id, _ := strconv.ParseInt(c.PostForm("product_id"), 10, 64)

	// Check user was paid product .
	// Have product
	var queryOrderHistory model.OrderHistory
	if err := db.GetDB().Where("username = ? AND product_id = ?", c.GetString("jwt_username"), id).First(&queryOrderHistory).Error; err == nil {

		// Create rate product , star point
		productRate := model.ProductRate{}
		productRate.ProductID = id
		productRate.Username = c.GetString("jwt_username")
		productRate.Rate = rate
		productRate.CreatedAt = time.Now()
		db.GetDB().Create(&productRate)

		// Get current product by Id
		var queryProduct model.Product
		db.GetDB().Where("id = ?", id).First(&queryProduct)

		// Update Rate Number total And Rate Star total .
		var product model.Product
		product.ID = uint(id)
		product.RateNum = queryProduct.RateNum + 1
		product.Rate = queryProduct.Rate + rate
		db.GetDB().Updates(&product)

		c.JSON(http.StatusOK, gin.H{"result": product})
	} else {
		// Not was bought product .
		c.JSON(http.StatusBadRequest, 0)
	}

}
