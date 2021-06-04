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

// SetupAddressAPI - Address for products delivery
func SetupAddressAPI(router *gin.Engine) {
	addressAPI := router.Group("/api")
	{
		// Get Address user owner list .
		addressAPI.GET("/address", interceptor.JwtVerify, getAddress)

		// Get Address by Id
		addressAPI.GET("/address/:id", interceptor.JwtVerify, getAddressByID)

		// Create Address
		addressAPI.POST("/address", interceptor.JwtVerify, createAddress)

		// Update Address .
		addressAPI.PUT("/address", interceptor.JwtVerify, editAddress)

		// Delete Address .
		addressAPI.DELETE("/address/:id", interceptor.JwtVerify, deleteAddress)
	}
}

func getAddress(c *gin.Context) {
	var address []model.Address

	keyword := c.Query("keyword")
	if keyword != "" {
		keyword = fmt.Sprintf("%%%s%%", keyword)
		db.GetDB().Where("name like ?", keyword).Find(&address)
	} else {
		db.GetDB().Find(&address).Where("username = ?", c.GetString("jwt_username")).Order("created_at desc")
	}
	c.JSON(http.StatusOK, address)

}

func getAddressByID(c *gin.Context) {
	var address model.Address
	db.GetDB().Where("id = ?", c.Param("id")).First(&address)
	c.JSON(http.StatusOK, address)
}

func createAddress(c *gin.Context) {
	address := model.Address{}
	address.Fname = c.PostForm("fname")
	address.Lname = c.PostForm("lname")
	address.Tel = c.PostForm("tel")
	address.AddressDes = c.PostForm("addressDes")
	address.Subdistric = c.PostForm("subdistric")
	address.Distric = c.PostForm("distric")
	address.Province = c.PostForm("province")
	address.Zipcode = c.PostForm("zipcode")
	address.UserID = c.PostForm("UserID")
	address.Default, _ = strconv.ParseBool(c.PostForm("default"))
	address.CreatedAt = time.Now()
	db.GetDB().Create(&address)

	c.JSON(http.StatusOK, gin.H{"result": address})

}
func editAddress(c *gin.Context) {
	address := model.Address{}
	id, _ := strconv.ParseInt(c.PostForm("id"), 10, 32)
	address.ID = uint(id)
	address.Fname = c.PostForm("fname")
	address.Lname = c.PostForm("lname")
	address.Tel = c.PostForm("tel")
	address.AddressDes = c.PostForm("addressDes")
	address.Subdistric = c.PostForm("subdistric")
	address.Distric = c.PostForm("distric")
	address.Province = c.PostForm("province")
	address.Zipcode = c.PostForm("zipcode")
	address.UserID = c.PostForm("UserID")
	address.Default, _ = strconv.ParseBool(c.PostForm("default"))
	address.CreatedAt = time.Now()
	db.GetDB().Updates(&address)

	c.JSON(http.StatusOK, gin.H{"result": address})
}

func deleteAddress(c *gin.Context) {

	id, _ := strconv.ParseInt(c.Param("id"), 10, 32)
	db.GetDB().Delete(&model.Address{}, id)
	c.JSON(http.StatusOK, gin.H{"result": "ok"})
}
