package main

import (
	"book-frontend-api-final/api"
	"fmt"
	"os"
	"time"

	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

func main() {
	router := gin.Default()

	// Set helmet
	router.Use(helmet.Default())
	router.Use(helmet.NoCache())

	// Set up CORS middleware options
	config := cors.Config{
		Origins:        "*",
		RequestHeaders: "Origin, Authorization, Content-Type",

		Methods:         "GET, POST, PUT, DELETE",
		Credentials:     true,
		ValidateHeaders: false,
		MaxAge:          1 * time.Minute,
	}

	// Apply the middleware to the router (works on groups too)
	router.Use(cors.Middleware(config))

	router.Static("/images", "./uploaded/images")

	api.Setup(router)
	// Start Set Port
	// router.Run(":8080")

	// in case on heroku
	var port = os.Getenv("PORT")
	if port == "" {
		fmt.Println("No Port In Heroku")
		router.Run()
	} else {
		fmt.Println("Environment Port : " + port)
		router.Run(fmt.Sprintf(":%s", port))
	}
}
