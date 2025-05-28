package routes

import (
	"example.com/rest-api/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventsById)

	authenticated := server.Group("/")
	authenticated.Use(middleware.Authenticate) // applying the authentication middleware to this group
	authenticated.POST("/events", createEvents)
	authenticated.PUT("/events/:id", updateEvents)
	authenticated.DELETE("/events/:id", deleteEvents)

	//server.POST("/events", middleware.Authenticate, createEvents)
	//server.PUT("/events/:id", updateEvents)
	//server.DELETE("/events/:id", deleteEvents)
	server.POST("/signup", signup)
	server.GET("/users", getUsers)
	server.POST("/login", login)
}
