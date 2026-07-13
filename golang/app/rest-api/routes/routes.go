package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events/:id", getEvent)       //GET request is used to get a specific event from the server
	server.GET("/events", getEvents)          //GET request is used to get the events from the server
	server.POST("/events", createEvents)      //POST request is used to create a new event
	server.PUT("/events/:id", updateEvent)    //PUT request is used to update a specific event from the server
	server.DELETE("/events/:id", deleteEvent) //DELETE request is used to delete a specific event from the server
	// server.POST("/signup")
}
