// main.go
// Author: Vaibhav Parmar

package main

import (
	"net/http"
	"event-booking-api/db"
	"event-booking-api/models"
	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()						// Initialize database connection
	server := gin.Default()			// Initialize Gin server

	//******** Routes ********//	
	
	server.GET("/", getRoot)
	server.GET("/health", healthCheck)
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	//*************************//

	server.Run(":8080")				// Start server on port 8080
}

// Handler for root endpoint
func getRoot(context *gin.Context) {
	//sending JSON response
	context.JSON(http.StatusOK, gin.H{
		"message": "Welcome to the Event Booking API",
		"author": "Vaibhav Parmar",
	})
}

// Handler for health check endpoint
func healthCheck(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"status": "healthy",
	})
}

// Handler to get all events
func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Could not fetch events": err.Error()})
		return
	}
	context.JSON(http.StatusOK, events)
}	

// Handler to create a new event
func createEvent(context *gin.Context) {
	var newEvent models.Event
	err := context.ShouldBindJSON(&newEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newEvent.ID = 1
	newEvent.UserID = 1
	err = models.AddEvent(newEvent)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Could not create a new event": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created successfully!", "event": newEvent})
}