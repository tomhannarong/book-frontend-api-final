package api

import (
	"book-frontend-api-final/db"
	"book-frontend-api-final/interceptor"
	"book-frontend-api-final/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupOrderHistoryAPI - Order History
func SetupOrderHistoryAPI(router *gin.Engine) {
	orderHistoryAPI := router.Group("/api")
	{
		// Get order history list form order master when successfuly paid
		orderHistoryAPI.GET("/orderHistory", interceptor.JwtVerify, getOrderHistory)

		// Get order history by Id
		orderHistoryAPI.GET("/orderHistory/:id", interceptor.JwtVerify, getOrderHistoryByID)
	}
}

func getOrderHistory(c *gin.Context) {
	var orderHistory []model.OrderHistory

	keyword := c.Query("keyword")
	if keyword != "" {
		keyword = fmt.Sprintf("%s", keyword)
		db.GetDB().Where("product_id = ?", keyword).Where("username = ?", c.GetString("jwt_username")).Find(&orderHistory)
	} else {
		db.GetDB().Where("username = ?", c.GetString("jwt_username")).Order("created_at desc").Find(&orderHistory)
	}
	c.JSON(http.StatusOK, orderHistory)

}

func getOrderHistoryByID(c *gin.Context) {
	orderHistory := model.OrderHistory{}
	db.GetDB().Where("id = ?", c.Param("id")).Where("username = ?", c.GetString("jwt_username")).First(&orderHistory)
	c.JSON(http.StatusOK, orderHistory)
}
