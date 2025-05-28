package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/rest-api/modells"
	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	//events := []string{"Event 1", "Event 2", "Event 3"}
	events, err := modells.GetEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "couldn't fetch events"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvents(context *gin.Context) {

	token := context.GetHeader("Authorization") // getting the token from the header of request
	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is required"})
		return
	}

	userIdfromToken, err := utils.VerfiyToken(token) // verifying the token
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		return
	}

	var event modells.Event
	if err = context.ShouldBindJSON(&event); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "coundn't parese data"})
		return
	}
	//event.ID = 1
	event.UserId = int(userIdfromToken) // setting the user ID from the token

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "couldn't save event"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created successfully", "event": event})
}

func getEventsById(context *gin.Context) {

	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}
	event, err := modells.GetEventById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "couldn't fetch event"})
		return
	}
	print(event)

	context.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Event %v with ID %v fetched successfully", event, id)})
}

func updateEvents(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	_, err = modells.GetEventById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "couldn't fetch event to update"})
		return
	}

	var updatedEvent modells.Event
	if err := context.ShouldBindJSON(&updatedEvent); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "coundn't parse data for incoming request of updating event"})
		return
	}
	updatedEvent.ID = id

	//err = updatedEvent.UpdateEvent()
	err = modells.UpdateEvent(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "couldn't update event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Event with ID %v updated successfully", id)})
}

func deleteEvents(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	err = modells.DeleteEvent(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "couldn't delete event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Event with ID %v deleted successfully", id)})
}
