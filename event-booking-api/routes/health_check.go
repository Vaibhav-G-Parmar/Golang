package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler for health check endpoint
func healthCheck(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"status": "healthy",
	})
}
