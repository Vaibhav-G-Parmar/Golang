// main.go
// Author: Vaibhav Parmar

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	// Initialize Gin server
	server := gin.Default()

	//******** Routes ********//	

	// Root endpoint
	server.GET("/", func(context *gin.Context) {
		//sending JSON response
		context.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the Event Booking API",
			"author": "Vaibhav Parmar",
		})
	})
	
	// Health check endpoint
	server.GET("/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"status": "healthy",
		})
	})
 
	// Start the server
	server.Run(":8080")
}