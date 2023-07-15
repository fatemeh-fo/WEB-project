// backend/handlers/auth.go

package handlers

import (
	"net/http"
	//"time"

	"github.com/gin-gonic/gin"

	"github.com/your_username/quiz-platform/backend/db"
	"github.com/your_username/quiz-platform/backend/models"
	"github.com/your_username/quiz-platform/backend/utils"

	"golang.org/x/crypto/bcrypt"
)

// LoginRequest represents the JSON request for the login endpoint.
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// SignupRequest represents the JSON request for the signup endpoint.
type SignupRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Login handles the user login process.
func Login(c *gin.Context) {
	var loginReq LoginRequest
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Retrieve the user from the database by username
	user, err := db.GetUserByUsername(loginReq.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Check if the provided password matches the hashed password in the database
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate JWT token
	token, err := utils.GenerateToken(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// Signup handles the user signup process.
func Signup(c *gin.Context) {
	var signupReq SignupRequest
	if err := c.ShouldBindJSON(&signupReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Check if the username already exists in the database
	_, err := db.GetUserByUsername(signupReq.Username)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
		return
	}

	// Hash the password before storing it in the database
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signupReq.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Create the user and store it in the database
	user := models.User{
		Username: signupReq.Username,
		Password: string(hashedPassword),
	}
	err = db.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

// ForgotPassword handles the user forgot password process.
func ForgotPassword(c *gin.Context) {
	// Implement the forgot password logic here
	// For example, you can send a password reset link to the user's email
	// and handle the password reset process separately.
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Forgot password functionality not implemented"})
}
