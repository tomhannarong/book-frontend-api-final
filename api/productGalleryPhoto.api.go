package api

import (
	"book-frontend-api-final/db"
	"book-frontend-api-final/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupProductGalleryPhotoAPI - Product Gallery Photo
func SetupProductGalleryPhotoAPI(router *gin.Engine) {
	productGalleryPhotoAPI := router.Group("/api")
	{
		// Get Product Gallery Photo by Id .
		productGalleryPhotoAPI.GET("/product/:id/gallery-photo" /*interceptor.JwtVerify,*/, getProductGalleryPhotoByID)
	}
}

func getProductGalleryPhotoByID(c *gin.Context) {
	productGalleryPhoto := model.ProductGalleryPhoto{}
	db.GetDB().Where("product_id = ?", c.Param("id")).First(&productGalleryPhoto)
	c.JSON(http.StatusOK, productGalleryPhoto)
}
