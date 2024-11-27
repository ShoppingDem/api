package controllers

import (
	"gin-user-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var newUser models.User

	// Bind incoming JSON to User struct and validate
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Assign an ID and add the new user to the slice
	newUser.ID = models.NextID
	models.NextID++
	models.Users = append(models.Users, newUser)

	// Respond with the created user
	c.JSON(http.StatusCreated, newUser)
}
