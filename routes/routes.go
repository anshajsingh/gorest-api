package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventsById)
	server.POST("/events", createEvents)
	server.PUT("/events/:id", updateEvents)
	server.DELETE("/events/:id", deleteEvents)
	server.POST("/signup", signup)
	server.GET("/users", getUsers)
	server.POST("/login", login)
}
