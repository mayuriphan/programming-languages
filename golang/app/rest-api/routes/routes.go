package routes

import (
	"rest-api/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events/:id", getEvent) //GET request is used to get a specific event from the server
	server.GET("/events", getEvents)

	auth := server.Group("/") //GET request is used to get the events from the server
	auth.Use(middlewares.Authenticate)
	auth.POST("/events", createEvents) //POST request is used to create a new event
	// server.POST("/events", middlewares.Authenticate, createEvents) //anoher way to add midldlware, left to right thus auth first then cretevent
	auth.POST("/events/:id/register", registerForEvent)
	auth.PUT("/events/:id", updateEvent)    //PUT request is used to update a specific event from the server
	auth.DELETE("/events/:id", deleteEvent) //DELETE request is used to delete a specific event from the server
	auth.DELETE("/events/:id/register", cancelRegistration)

	server.POST("/signup", signup) //POST request is used to sign up a new user
	server.POST("/login", login)
}
