package api

import (
	"book-frontend-api-final/db"
	"book-frontend-api-final/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupBoardPostAPI - WebBoard news list
func SetupBoardPostAPI(router *gin.Engine) {
	boardPostAPI := router.Group("/api")
	{
		// Get WebBoard news list
		boardPostAPI.GET("/boardPost" /*interceptor.JwtVerify,*/, getBoardPost)

		// Get WebBoard by Id
		boardPostAPI.GET("/boardPost/:id" /*interceptor.JwtVerify,*/, getBoardPostByID)
	}
}

func getBoardPost(c *gin.Context) {
	var boardPost []model.BoardPost

	keyword := c.Query("keyword")
	if keyword != "" {
		keyword = fmt.Sprintf("%%%s%%", keyword)
		db.GetDB().Where("topic like ?", keyword).Find(&boardPost)
	} else {
		db.GetDB().Find(&boardPost).Order("created_at desc")
	}
	c.JSON(http.StatusOK, boardPost)

}

func getBoardPostByID(c *gin.Context) {
	boardPost := model.BoardPost{}
	db.GetDB().Where("id = ?", c.Param("id")).First(&boardPost)
	c.JSON(http.StatusOK, boardPost)
}
