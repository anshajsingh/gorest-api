package main

import (
	"fmt"

	"example.com/rest-api/database"
	"example.com/rest-api/routes"
	gin "github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	if _, err := database.InitDB(); err != nil {
		panic("Couldn't connect to database: " + err.Error())
	}

	routes.RegisterRoutes(server)

	server.Run(":8080")
	fmt.Println("Server is running on port 8080")
}
