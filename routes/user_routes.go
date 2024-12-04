package routes

import (
	"github.com/ShoppingDem/api/controllers"
	"github.com/gin-gonic/gin"
)

// RegisterUserRoutes sets up the routes for user-related endpoints.
// It creates a group of routes under "/users" and registers routes for user operations.
//
// @Summary Register user routes
// @Description Sets up the routes for user-related endpoints
// @Tags users
// @Accept json
// @Produce json
// @Router /users [post]
func RegisterUserRoutes(r *gin.Engine) {
	userGroup := r.Group("/users")
	{
		userGroup.POST("/", controllers.CreateUser)
		userGroup.POST("/register", controllers.CreateUser)
	}
}
