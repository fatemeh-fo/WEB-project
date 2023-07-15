package handlers

import (
	"encoding/json"
	"net/http"

	
	"github.com/gin-gonic/gin"
	"github.com/fatemeh-fo/WEB-project/Back/models"
)

// CreateQuizHandler handles the creation of a new quiz.
func CreateQuizHandler(c *gin.Context) {
	var quiz models.Quiz

	// Parse the request body and decode it into the quiz struct
	if err := c.ShouldBindJSON(&quiz); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Call the CreateQuiz function to save the quiz to the database
	err := models.CreateQuiz(&quiz)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create quiz"})
		return
	}

	// Return the newly created quiz as the response
	c.JSON(http.StatusCreated, quiz)
}

// GetQuizByIDHandler retrieves a quiz from the database by its ID.
func GetQuizByIDHandler(c *gin.Context) {
	quizID := c.Param("id")

	// Call the GetQuizByID function to retrieve the quiz from the database by ID
	quiz, err := models.GetQuizByID(quizID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Quiz not found"})
		return
	}

	// Return the quiz as the response
	c.JSON(http.StatusOK, quiz)
}

// SubmitQuizHandler handles the submission of quiz answers.
func SubmitQuizHandler(c *gin.Context) {
	var answers map[string]string

	// Parse the request body and decode it into the answers map
	if err := json.NewDecoder(c.Request.Body).Decode(&answers); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	quizID := c.Param("id")

	// Call the GetQuizByID function to retrieve the quiz from the database by ID
	quiz, err := models.GetQuizByID(quizID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Quiz not found"})
		return
	}

	// Calculate the score based on the user's answers
	score := calculateScore(&quiz, answers)

	// Return the score as the response
	c.JSON(http.StatusOK, gin.H{"score": score})
}

// calculateScore calculates the score based on the user's answers and the correct answers in the quiz.
func calculateScore(quiz *models.Quiz, userAnswers map[string]string) int {
	// Implementation required: Calculate the score based on the correct answers
	// ...
	return 0
}
