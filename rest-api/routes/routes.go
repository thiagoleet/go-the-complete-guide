package routes

import (
	"github.com/gin-gonic/gin"
	"thiagoleite.me/go-event-booking-api/middlewares"
)

func RegisterRoutes(server *gin.Engine) {

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	autenticated := server.Group("/")
	autenticated.Use(middlewares.Authenticate)
	autenticated.POST("/events", createEvent)
	autenticated.PUT("/events/:id", updateEvent)
	autenticated.DELETE("/events/:id", deleteEvent)

	server.POST("/signup", signUp)
	server.POST("/login", login)
}
