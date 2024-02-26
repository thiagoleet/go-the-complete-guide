package main

import (
	"github.com/gin-gonic/gin"
	"thiagoleite.me/go-event-booking-api/db"
	"thiagoleite.me/go-event-booking-api/routes"
)

func main() {
	db.InitDB()

	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
