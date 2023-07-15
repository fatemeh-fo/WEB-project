// backend/main.go

package main

import (
	"log"
	"net/http"

	"github.com/fatemeh-fo/WEB-project//Back/handlers"
	"github.com/fatemeh-fo/WEB-project/Back/db/database.go"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database
	err := db.Initialize("postgres://username:password@localhost/dbname?sslmode=disable", "path/to/schema.sql")
	if err != nil {
		log.Fatal(err)
	}

	// Initialize the Gin router
	router := gin.Default()

	// Set up routes and handlers
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the Quiz Platform API",
		})
	})

	router.POST("/login", handlers.Login)
	router.POST("/signup", handlers.Signup)
	router.POST("/forgot-password", handlers.ForgotPassword)
	router.GET("/profile", handlers.GetProfile)
	router.PUT("/profile", handlers.UpdateProfile)
	router.POST("/quiz", handlers.CreateQuiz)
	router.GET("/quiz/:id", handlers.GetQuiz)
	router.POST("/quiz/:id/submit", handlers.SubmitQuiz)

	// Start the server
	err = router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
