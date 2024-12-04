package controllers

import (
	"net/http"

	"github.com/ShoppingDem/api/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// ErrorResponse represents the structure of an error response
type ErrorResponse struct {
	Error string `json:"error"`
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the input payload
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body models.UserRegistration true "Create user"
// @Success 201 {object} models.User
// @Failure 400 {object} ErrorResponse
// @Router /users [post]
func CreateUser(c *gin.Context) {
	var registration models.UserRegistration

	if err := c.ShouldBindJSON(&registration); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	if len(registration.PhoneNumber) < 10 {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid phone number"})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registration.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to hash password"})
		return
	}

	newUser := models.User{
		ID:           models.NextID,
		PhoneNumber:  registration.PhoneNumber,
		PasswordHash: string(hashedPassword),
	}

	models.NextID++
	models.Users = append(models.Users, newUser)

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully", "userId": newUser.ID})
}
