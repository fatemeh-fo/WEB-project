package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	//TODO :  Change this to your own secret key for JWT
	jwtSecret = "your-secret-key"
)

type Question struct {
	ID         int64    `json:"id"`
	Text       string   `json:"text"`
	Choices    []string `json:"choices"`
	CorrectAns string   `json:"correct_answer"`
}

type Quiz struct {
	ID        int64      `json:"id"`
	Title     string     `json:"title"`
	Questions []Question `json:"questions"`
}

var db *sql.DB
var quizzes []Quiz

func main() {
	var err error
	db, err = sql.Open("postgres", "postgres://username:password@localhost/dbname?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := gin.Default()

	// Set up session middleware
	store := cookie.NewStore([]byte("secret-key"))
	router.Use(sessions.Sessions("session-name", store))

	router.LoadHTMLGlob("templates/*")

	// Render the quiz creation form
	router.GET("/create", func(c *gin.Context) {
		c.HTML(http.StatusOK, "create.html", nil)
	})

	// Handle quiz creation form submission
	router.POST("/create", func(c *gin.Context) {
		// Extract quiz data from the request
		title := c.PostForm("title")
		questionText := c.PostForm("question")
		choices := c.PostFormArray("choices")
		correctAns := c.PostForm("correct_answer")

		// Create a new question
		question := Question{
			Text:       questionText,
			Choices:    choices,
			CorrectAns: correctAns,
		}

		// Create a new quiz
		quiz := Quiz{
			Title:     title,
			Questions: []Question{question},
		}

		// Save the quiz to the database and get the assigned ID
		quizID, err := saveQuizToDatabase(quiz)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save quiz"})
			return
		}
		quiz.ID = quizID

		// Add the quiz to the list of quizzes (in-memory storage)
		quizzes = append(quizzes, quiz)

		// Redirect to a success page or quiz list page
		c.Redirect(http.StatusSeeOther, "/quizzes")
	})

	// Render the list of quizzes
	router.GET("/quizzes", func(c *gin.Context) {
		quizzes, err := getQuizzesFromDatabase()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve quizzes"})
			return
		}

		c.HTML(http.StatusOK, "quizzes.html", gin.H{
			"quizzes": quizzes,
		})
	})

	// Render the quiz taking page
	router.GET("/quiz/:id", func(c *gin.Context) {
		quizID := c.Param("id")

		// Retrieve the quiz from the database by its ID
		quiz, err := getQuizFromDatabase(quizID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve quiz"})
			return
		}

		c.HTML(http.StatusOK, "quiz.html", gin.H{
			"quiz": quiz,
		})
	})

	// Handle quiz submission
	router.POST("/quiz/:id/submit", func(c *gin.Context) {
		quizID := c.Param("id")

		// Retrieve user's answers from the request
		answers := c.PostFormMap("answers")

		// Retrieve the quiz from the database by its ID
		quiz, err := getQuizFromDatabase(quizID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve quiz"})
			return
		}

		// Calculate the score
		score := calculateScore(quiz, answers)

		// Render the score page
		c.HTML(http.StatusOK, "score.html", gin.H{
			"score": score,
		})
	})

	// Render the user registration form
	router.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", nil)
	})

	// Handle user registration form submission
	router.POST("/register", func(c *gin.Context) {
		// Extract user registration data from the request
		username := c.PostForm("username")
		password := c.PostForm("password")

		// Perform user registration (implementation required)
		err := registerUser(username, password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
			return
		}

		// Redirect to a success page or login page
		c.Redirect(http.StatusSeeOther, "/login")
	})

	// Render the user login form
	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	// Handle user login form submission
	router.POST("/login", func(c *gin.Context) {
		// Extract user login credentials from the request
		username := c.PostForm("username")
		password := c.PostForm("password")

		// Perform user authentication (implementation required)
		valid, err := authenticateUser(username, password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to authenticate user"})
			return
		}
		if !valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}

		// Generate JWT token
		token, err := generateToken(username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}

		// Set token in the session
		session := sessions.Default(c)
		session.Set("token", token)
		session.Save()

		// Redirect to a dashboard or home page
		c.Redirect(http.StatusSeeOther, "/dashboard")
	})

	// Protected route - require authentication
	router.GET("/dashboard", authMiddleware(), func(c *gin.Context) {
		// Get user information from the token
		token := c.MustGet("token").(*jwt.Token)
		claims := token.Claims.(jwt.MapClaims)
		username := claims["username"].(string)

		c.HTML(http.StatusOK, "dashboard.html", gin.H{
			"username": username,
		})
	})

	// Logout route
	router.GET("/logout", func(c *gin.Context) {
		// Clear session and redirect to login page
		session := sessions.Default(c)
		session.Clear()
		session.Save()
		c.Redirect(http.StatusSeeOther, "/login")
	})

	err = router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

// Function to save the quiz to the database and return the assigned ID
func saveQuizToDatabase(quiz Quiz) (int64, error) {
	// Implementation required: Insert the quiz into the database and retrieve the assigned ID
	// ...
	return 1, nil
}

// Function to retrieve quizzes from the database
func getQuizzesFromDatabase() ([]Quiz, error) {
	// Implementation required: Retrieve quizzes from the database
	// ...
	return quizzes, nil
}

// Function to retrieve a quiz from the database by its ID
func getQuizFromDatabase(quizID string) (Quiz, error) {
	// Implementation required: Retrieve the quiz from the database by its ID
	// ...
	return Quiz{}, nil
}

// Function to calculate the score based on the user's answers
func calculateScore(quiz Quiz, userAnswers map[string]string) int {
	// Implementation required: Calculate the score based on the correct answers
	// ...
	return 0
}

// Function to register a user (dummy implementation)
func registerUser(username, password string) error {
	// Implementation required: Register the user in the database
	// ...
	return nil
}

// Function to authenticate a user (dummy implementation)
func authenticateUser(username, password string) (bool, error) {
	// Implementation required: Authenticate the user from the database
	// ...
	return true, nil
}

// Function to generate JWT token
func generateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	})

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Middleware to handle authentication
func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get token from the session
		session := sessions.Default(c)
		token := session.Get("token")

		// Check if token exists and is valid
		if token == nil {
			c.Redirect(http.StatusSeeOther, "/login")
			c.Abort()
			return
		}

		tokenString := fmt.Sprintf("%v", token)
		parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret), nil
		})

		if err != nil {
			c.Redirect(http.StatusSeeOther, "/login")
			c.Abort()
			return
		}

		if parsedToken.Valid {
			c.Set("token", parsedToken)
			c.Next()
		} else {
			c.Redirect(http.StatusSeeOther, "/login")
			c.Abort()
		}
	}
}
