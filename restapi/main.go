package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"restapi/api"
)

func main() {
	// Initialize the Gin router
	router := gin.Default()

	// Setup API routes
	apiRouter := api.SetupRouter()

	// Convert the apiRouter.ServeHTTP function into a gin.HandlerFunc
	apiHandler := gin.HandlerFunc(func(c *gin.Context) {
		apiRouter.ServeHTTP(c.Writer, c.Request)
	})

	// Add the API router as middleware
	router.Use(apiHandler)

	// Run the server on port 8080
	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
