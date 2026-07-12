package main

import (
	"net/http"
	"rest-api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()              //server is the engine pointer
	server.GET("/events", getEvents)     //GET request is used to get the events from the server
	server.POST("/events", createEvents) //POST request is used to create a new event
	server.Run(":8080")                  // listen and serve on localhost:8080
}

func getEvents(context *gin.Context) { //param is send using context pointer
	// Logic to retrieve events from the database or any other source
	events := models.GetAllEvents() //call the function from models package to get all events

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
	event.Save()
	event.ID = 1
	event.UserID = 001
	context.JSON(http.StatusCreated, gin.H{"message": "event created", "event": event}) //call the function from models package to save the event
}
