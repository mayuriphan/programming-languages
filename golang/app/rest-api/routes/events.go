package routes

import (
	"net/http"
	"rest-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadGateway, gin.H{"error": "Event not found"})
		return
	}

	event, err := models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Event not found"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"event": event})
}

func getEvents(context *gin.Context) { //param is send using context pointer
	// Logic to retrieve events from the database or any other source
	events, err := models.GetAllEvents() //call the function from models package to get all events
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{ //map type is used to send the response in JSON format
		"events": events,
	})
}

func createEvents(context *gin.Context) {
	var event models.Event                       //create a new event variable of type Event
	err := context.Copy().ShouldBindJSON(&event) // act as scanner in fmt package
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = event.Save() //call the function from models package to save the event
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "event created", "event": event}) //call the function from models package to save the event
}

func updateEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}
	_, err = models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Event not found"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedEvent.ID = eventID
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "event updated successfully !", "event": updatedEvent})
}

func deleteEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}
	event, err := models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Event not found"})
		return
	}
	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Event not found"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "event deleted successfully !"})
}
