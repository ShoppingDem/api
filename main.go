package main

import (
	"gin-user-app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Initialize routes
	routes.RegisterUserRoutes(r)

	r.Run(":8080") // Start server on localhost:8080
}
