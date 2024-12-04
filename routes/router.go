package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/ShoppingDem/api/docs"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Serve static files
	r.Static("/static", "./static")

	// Serve the home page
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{})
	})

	// Protected route
	r.GET("/api/protected", func(c *gin.Context) {
		claims, exists := c.Get("claims")
		if !exists {
			c.JSON(401, gin.H{"error": "No claims found"})
			return
		}
		c.JSON(200, gin.H{"message": "This is a protected route", "claims": claims})
	})

	// Swagger documentation route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
