package models

// Question represents a question in a quiz.
type Question struct {
	ID         int      `json:"id"`
	Text       string   `json:"text"`
	Choices    []string `json:"choices"`
	CorrectAns string   `json:"correct_answer"`
}

// CreateQuestion creates a new question in the database.
func CreateQuestion(question *Question) error {
	// Implement the logic to save the question to the database
	// ...
	return nil
}

// GetQuestionByID retrieves a question from the database by ID.
func GetQuestionByID(questionID int) (*Question, error) {
	// Implement the logic to retrieve the question from the database by ID
	// ...
	return nil, nil
}
