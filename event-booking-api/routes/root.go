package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler for root endpoint
func getRoot(context *gin.Context) {
	//sending JSON response
	context.JSON(http.StatusOK, gin.H{
		"message": "Welcome to the Event Booking API",
		"author":  "Vaibhav Parmar",
	})
}
