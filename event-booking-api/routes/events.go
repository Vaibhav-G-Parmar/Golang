package routes

import (
	"event-booking-api/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// handler for events related services

// get all events
func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Could not fetch events": err.Error()})
		return
	}
	context.JSON(http.StatusOK, events)
}

// Handler to get event by ID
func getEventById(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch event"})
		return
	}
	if event == nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}
	context.JSON(http.StatusOK, event)
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

// Handler to update an existing event
func updateEvent(context *gin.Context) {

	// inside handler
	log.Printf("Updating event with ID: %s, client=%s, method=%s",
		context.Param("id"), context.ClientIP(), context.Request.Method)

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse event id"})
		return
	}

	_, err = models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch event"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		// Log detailed binding error for debugging and return a useful message to the caller
		log.Printf("Failed to bind updated event JSON for id=%d: %v", eventId, err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse updated event data", "details": err.Error()})
		return
	}

	log.Printf("Updating event ID: %d with new data: %+v", eventId, updatedEvent)

	updatedEvent.ID = eventId
	err = models.UpdateEvent(updatedEvent)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully!", "event_id": eventId, "updated_event": updatedEvent})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse event id"})
		return
	}

	_, err = models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch the event"})
		return
	}

	err = models.DeleteEvent(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully!", "event_id": eventId})
}

