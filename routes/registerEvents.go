package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/modells"
	"github.com/gin-gonic/gin"
)

func registerEvent(context *gin.Context) {

	userIdfromToken := context.GetInt64("userId") // getting the user ID from the context set by the middleware
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid event ID"})
		return
	}

	event, err := modells.GetEventById(id) // Fetch the event to ensure it exists
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "couldn't fetch event"})
		return
	}

	err = event.RegisterEvent(userIdfromToken)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "couldn't register for event"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Registration successful", "event": event})

}

func unregisterEvent(context *gin.Context) {
	userIdfromToken := context.GetInt64("userId") // getting the user ID from the context set by the middleware
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid event ID"})
		return
	}

	event, err := modells.GetEventById(id) // Fetch the event to ensure it exists
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "couldn't fetch event"})
		return
	}

	err = event.UnregisterEvent(userIdfromToken)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "couldn't unregister from event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Unregistration successful", "event": event})
}
