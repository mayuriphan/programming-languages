package main

import (
	"rest-api/db"
	"rest-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB() //initialize the database connection

	server := gin.Default() //server is the engine pointer

	routes.RegisterRoutes(server) //register the routes

	server.Run(":8080") // listen and serve on localhost:8080
}
