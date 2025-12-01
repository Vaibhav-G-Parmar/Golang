package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	// Initialize Gin server
	server := gin.Default()

	//******** Routes ********//	

	// Root endpoint
	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the Event Booking API",
			"author": "Vaibhav Parmar",
		})
	})
	
	// Health check endpoint
	server.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "healthy",
		})
	})

	// Start the server
	server.Run(":8080")
}