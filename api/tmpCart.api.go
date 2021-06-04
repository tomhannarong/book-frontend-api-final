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

// SetupTmpCartAPI - Cart of user customer
func SetupTmpCartAPI(router *gin.Engine) {
	tmpCartAPI := router.Group("/api")
	{
		// Get all product in cart
		tmpCartAPI.GET("/cart", interceptor.JwtVerify, getTmpCart)

		// Create cart when first step buy product
		tmpCartAPI.POST("/cart", interceptor.JwtVerify, createTmpCart)

		// Delete all product in cart
		tmpCartAPI.DELETE("/cart/:id", interceptor.JwtVerify, deleteTmpCart)

		// Add quantity product in cart
		tmpCartAPI.POST("/cart/addQuaProduct", interceptor.JwtVerify, addQuaTmpCart)

		// Remove quantity product in cart
		tmpCartAPI.DELETE("/cart/removeQuaProduct/:id", interceptor.JwtVerify, removeQuaTmpCart)

		// Delete product in cart
		tmpCartAPI.DELETE("/cart/deleteProduct/:id", interceptor.JwtVerify, deleteProductTmpCart)
	}
}

func getTmpCart(c *gin.Context) {
	var tmpCart []model.TmpCart

	keyword := c.Query("keyword")
	if keyword != "" {
		keyword = fmt.Sprintf("%s", keyword)
		db.GetDB().Where("book_id = ?", keyword).Where("username = ?", c.GetString("jwt_username")).Find(&tmpCart)
	} else {
		db.GetDB().Find(&tmpCart).Where("username = ?", c.GetString("jwt_username")).Order("created_at desc")
	}
	c.JSON(http.StatusOK, tmpCart)

}

func createTmpCart(c *gin.Context) {

	tmpCart := model.TmpCart{}
	if c.ShouldBind(&tmpCart) == nil {
		BookID, _ := strconv.ParseInt(c.PostForm("book_id"), 10, 64)
		quantity, _ := strconv.ParseInt(c.PostForm("quantity"), 10, 64)
		UserID, _ := strconv.ParseInt(c.PostForm("user_id"), 10, 64)

		tmpCart.BookID = BookID
		tmpCart.Username = c.PostForm("username")
		tmpCart.Quantity = quantity
		tmpCart.IsblameProduct = c.PostForm("blame_product")
		tmpCart.IsBuffet, _ = strconv.ParseBool(c.PostForm("buffet"))
		tmpCart.IsDiscount, _ = strconv.ParseBool(c.PostForm("can_discount"))
		tmpCart.UserID = UserID
		tmpCart.IsEbook, _ = strconv.ParseBool(c.PostForm("is_ebook"))
		tmpCart.CreatedAt = time.Now()
		db.GetDB().Create(&tmpCart)

		c.JSON(http.StatusOK, gin.H{"result": tmpCart})
	} else {
		c.JSON(401, gin.H{"status": "unable to bind data"})
	}

}

func addQuaTmpCart(c *gin.Context) {

	tmpCart := model.TmpCart{}
	if c.ShouldBind(&tmpCart) == nil {
		var queryTmpCart model.TmpCart

		BookID, _ := strconv.ParseInt(c.PostForm("book_id"), 10, 64)
		UserID, _ := strconv.ParseInt(c.PostForm("user_id"), 10, 64)

		tmpCart.BookID = BookID
		tmpCart.Username = c.PostForm("username")
		tmpCart.IsblameProduct = c.PostForm("blame_product")
		tmpCart.IsBuffet, _ = strconv.ParseBool(c.PostForm("buffet"))
		tmpCart.IsDiscount, _ = strconv.ParseBool(c.PostForm("can_discount"))
		tmpCart.UserID = UserID
		tmpCart.IsEbook, _ = strconv.ParseBool(c.PostForm("is_ebook"))
		tmpCart.CreatedAt = time.Now()

		// Check , Have a product in cart
		// 1 . if Have Update product in cart .
		// 2 . if Not have Create new product in cart .
		if err := db.GetDB().First(&queryTmpCart, "book_id = ? AND username = ?", tmpCart.BookID, c.GetString("jwt_username")).Error; err != nil {
			tmpCart.ID = uint(queryTmpCart.ID)
			tmpCart.Quantity = queryTmpCart.Quantity + 1
			db.GetDB().Updates(&tmpCart)
		} else {
			tmpCart.Quantity = 1
			db.GetDB().Create(&tmpCart)
		}
		c.JSON(http.StatusOK, gin.H{"result": tmpCart})
	} else {
		c.JSON(401, gin.H{"status": "unable to bind data"})
	}

}

func removeQuaTmpCart(c *gin.Context) {

	tmpCart := model.TmpCart{}
	if c.ShouldBind(&tmpCart) == nil {
		var queryTmpCart model.TmpCart

		BookID, _ := strconv.ParseInt(c.PostForm("book_id"), 10, 64)
		UserID, _ := strconv.ParseInt(c.PostForm("user_id"), 10, 64)

		tmpCart.BookID = BookID
		tmpCart.Username = c.PostForm("username")
		tmpCart.Quantity = queryTmpCart.Quantity - 1
		tmpCart.IsblameProduct = c.PostForm("blame_product")
		tmpCart.IsBuffet, _ = strconv.ParseBool(c.PostForm("buffet"))
		tmpCart.IsDiscount, _ = strconv.ParseBool(c.PostForm("can_discount"))
		tmpCart.UserID = UserID
		tmpCart.IsEbook, _ = strconv.ParseBool(c.PostForm("is_ebook"))
		tmpCart.CreatedAt = time.Now()

		// Check product Quantity in cart ( == 1 ) is delete products in cart
		if queryTmpCart.Quantity == 1 {
			// Delete product in cart .
			db.GetDB().Delete(&model.TmpCart{}, BookID)
		} else {
			// Update product ( Quantity - 1 )
			db.GetDB().Updates(&tmpCart)
		}

		c.JSON(http.StatusOK, gin.H{"result": tmpCart})
	} else {
		c.JSON(401, gin.H{"status": "unable to bind data"})
	}

}

func deleteProductTmpCart(c *gin.Context) {
	// Get Id product from param
	id, _ := strconv.ParseInt(c.Param("id"), 10, 32)
	db.GetDB().Where("book_id = ? AND username = ?", id, c.GetString("jwt_username")).Delete(&model.TmpCart{})
	c.JSON(http.StatusOK, gin.H{"result": "ok"})

}

func deleteTmpCart(c *gin.Context) {
	// Get Id Temp Cart from param
	id, _ := strconv.ParseInt(c.Param("id"), 10, 32)
	db.GetDB().Delete(&model.TmpCart{}, id)
	c.JSON(http.StatusOK, gin.H{"result": "ok"})
}
