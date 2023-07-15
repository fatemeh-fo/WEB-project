package handlers

import (
	//"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/your-username/quiz-platform/backend/models"
	"github.com/your_username/quiz-platform/backend/utils"
)

// RegisterUserHandler handles user registration.
func RegisterUserHandler(c *gin.Context) {
	var user models.User

	// Parse the request body and decode it into the user struct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Call the CreateUser function to save the user to the database
	err := models.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	// Return a success response
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// LoginUserHandler handles user login.
func LoginUserHandler(c *gin.Context) {
	var user models.User

	// Parse the request body and decode it into the user struct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Call the GetUserByUsername function to retrieve the user from the database
	existingUser, err := models.GetUserByUsername(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to authenticate user"})
		return
	}

	// Check if the provided password matches the stored password
	if existingUser.Password != user.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate JWT token
	token, err := utils.GenerateToken(existingUser.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Return the JWT token as the response
	c.JSON(http.StatusOK, gin.H{"token": token})
}
