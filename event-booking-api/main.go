// main.go
// Author: Vaibhav Parmar

package main

import (
	"event-booking-api/db"
	"event-booking-api/routes"
	
	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()						// Initialize database connection
	server := gin.Default()			// Initialize Gin server
	routes.RegisterRoutes(server)	// Register API routes
	server.Run(":8080")				// Start server on port 8080
}

