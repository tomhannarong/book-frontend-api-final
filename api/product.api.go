package api

import (
	"book-frontend-api-final/db"
	"book-frontend-api-final/interceptor"
	"book-frontend-api-final/model"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// SetupProductAPI - all Products [ book + e-book ]
func SetupProductAPI(router *gin.Engine) {
	productAPI := router.Group("/api/")
	{
		// Get all Products own list [ book + ebook ]
		productAPI.GET("/product" /*interceptor.JwtVerify,*/, getProduct)

		// Get product by Id
		productAPI.GET("/product/:id" /*interceptor.JwtVerify,*/, getProductByID)

		// Create Product own by user or Publisher own
		productAPI.POST("/product", interceptor.JwtVerify, createProduct)

		// Update Product own by user or Publisher own
		productAPI.PUT("/product", interceptor.JwtVerify, editProduct)

		// Delete Product own by user or Publisher own
		productAPI.DELETE("/product/:id", interceptor.JwtVerify, deleteProduct)
	}
}

/*
func getProduct(c *gin.Context) {
	var product []model.Product
	db.GetDB().Find(&product)
	c.JSON(http.StatusOK, product)
}
*/
func getProduct(c *gin.Context) {
	var product []model.Product

	keyword := c.Query("keyword")
	if keyword != "" {
		keyword = fmt.Sprintf("%%%s%%", keyword)
		db.GetDB().Where("book_name like ?", keyword).Find(&product)
	} else {
		db.GetDB().Find(&product)
	}
	c.JSON(http.StatusOK, product)

}

func getProductByID(c *gin.Context) {
	var product model.Product
	db.GetDB().Where("id = ?", c.Param("id")).First(&product)
	c.JSON(http.StatusOK, product)
}

func createProduct(c *gin.Context) {
	product := model.Product{}
	// product.Name = c.PostForm("name")
	// product.Stock, _ = strconv.ParseInt(c.PostForm("stock"), 10, 64)
	product.Price, _ = strconv.ParseFloat(c.PostForm("price"), 64)
	product.CreatedAt = time.Now()
	db.GetDB().Create(&product)
	image, _ := c.FormFile("image")

	saveImage(image, &product, c)
	c.JSON(http.StatusOK, gin.H{"result": product})

}

func editProduct(c *gin.Context) {
	var product model.Product
	id, _ := strconv.ParseInt(c.PostForm("id"), 10, 32)
	product.ID = uint(id)
	// product.Name = c.PostForm("name")
	// product.Stock, _ = strconv.ParseInt(c.PostForm("stock"), 10, 64)
	product.Price, _ = strconv.ParseFloat(c.PostForm("price"), 64)
	product.CreatedAt = time.Now()
	db.GetDB().Updates(&product)

	image, _ := c.FormFile("image")
	saveImage(image, &product, c)
	c.JSON(http.StatusOK, gin.H{"result": product})

}

func deleteProduct(c *gin.Context) {

	id, _ := strconv.ParseInt(c.Param("id"), 10, 32)
	db.GetDB().Delete(&model.Product{}, id)
	c.JSON(http.StatusOK, gin.H{"result": "ok"})
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func saveImage(image *multipart.FileHeader, product *model.Product, c *gin.Context) {
	if image != nil {
		runningDir, _ := os.Getwd()
		// product.Image = image.Filename
		extension := filepath.Ext(image.Filename)
		fileName := fmt.Sprintf("%d%s", product.ID, extension)
		filePath := fmt.Sprintf("%s/uploaded/images/%s", runningDir, fileName)

		if fileExists(filePath) {
			os.Remove(filePath)
		}
		c.SaveUploadedFile(image, filePath)
		db.GetDB().Model(&product).Update("image", fileName)
	}
}
