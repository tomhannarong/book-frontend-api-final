package api

import (
	"book-frontend-api-final/db"
	"book-frontend-api-final/interceptor"
	"book-frontend-api-final/model"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// SetupOrderMasAPI - Order Master .
func SetupOrderMasAPI(router *gin.Engine) {
	orderMasAPI := router.Group("/api")
	{
		// Get order master for user customer  , Read only
		orderMasAPI.GET("/order", interceptor.JwtVerify, getOrderMas)

		// Get order master by id
		orderMasAPI.GET("/order/:id", interceptor.JwtVerify, getOrderMasByID)

		// Create or Add order master from all products item in cart of user customer
		// Create Tran Master at list all product item of order master
		orderMasAPI.POST("/order", interceptor.JwtVerify, createOrderMas)

		// Delete or Cancel order master , already input reason
		orderMasAPI.DELETE("/order/:id", interceptor.JwtVerify, deleteOrderMas)
	}
}

func getOrderMas(c *gin.Context) {
	var orderMas []model.OrderMas

	keyword := c.Query("keyword")
	if keyword != "" {
		keyword = fmt.Sprintf("%s", keyword)
		db.GetDB().Where("order_id = ?", keyword).Where("username = ?", c.GetString("jwt_username")).Find(&orderMas)
	} else {
		db.GetDB().Where("username = ?", c.GetString("jwt_username")).Order("created_at desc").Find(&orderMas)
	}
	c.JSON(http.StatusOK, orderMas)

}

func getOrderMasByID(c *gin.Context) {
	orderMas := model.OrderMas{}
	db.GetDB().Where("id = ?", c.Param("id")).Where("username = ?", c.GetString("jwt_username")).First(&orderMas)
	c.JSON(http.StatusOK, orderMas)
}

func createOrderMas(c *gin.Context) {

	// Get Product in Cart
	var queryTmpCart []model.TmpCart
	db.GetDB().Where("username = ?", c.GetString("jwt_username")).Find(&queryTmpCart)

	// Loop Product in cart .
	var net float64 = 0.0
	for _, item := range queryTmpCart {

		// Create Order Tran
		orderTran := model.OrderTran{}

		orderTran.BookID = item.BookID
		orderTran.Username = item.Username
		orderTran.Quantitys = item.Quantity
		orderTran.ProductPrice = item.Product.ProductPrice
		orderTran.IsEbook = item.IsEbook
		orderTran.CreatedAt = time.Now()

		db.GetDB().Create(&orderTran)

		// Sum price product
		price := item.Product.ProductPrice
		net += price
	}

	// Delete all product in Cart
	db.GetDB().Where("username = ?", c.GetString("jwt_username")).Delete(&model.TmpCart{})

	AddressID, _ := strconv.ParseInt(c.PostForm("address_id"), 10, 64)
	UserID, _ := strconv.ParseInt(c.GetString("jwt_staff_id"), 10, 64)
	transportRate, _ := strconv.ParseFloat(c.GetString("transport_rate"), 2)

	// Create Order Master
	orderMas := model.OrderMas{}
	orderMas.Username = c.PostForm("Username")
	orderMas.Payment = c.PostForm("Payment")
	orderMas.Transport = c.PostForm("Transport")
	orderMas.TransportRate = transportRate
	orderMas.NetPrice = net
	orderMas.OrderStatus = "y"
	orderMas.ShowStatus = c.PostForm("show_status")
	orderMas.TranferAddress = c.PostForm("tranfer_address")
	orderMas.AddressID = AddressID
	orderMas.UserID = UserID
	orderMas.CreatedAt = time.Now()
	db.GetDB().Create(&orderMas)

	c.JSON(http.StatusOK, gin.H{"result": orderMas})

}

func deleteOrderMas(c *gin.Context) {

	id, _ := strconv.ParseInt(c.Param("id"), 10, 32)

	db.GetDB().Where("username = ?", c.GetString("jwt_username")).Delete(&model.OrderMas{}, id)

	db.GetDB().Where("username = ?", c.GetString("jwt_username")).Where("order_id = ?", id).Delete(&model.OrderTran{})

	c.JSON(http.StatusOK, gin.H{"result": "ok"})
}
