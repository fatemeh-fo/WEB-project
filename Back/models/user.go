// backend/models/user.go

package models

import "time"

// User represents a user in the system.
type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CreateUser creates a new user in the database.
func CreateUser(user *User) error {
	// Implement the logic to save the user to the database
	// ...
	return nil
}

// GetUserByID retrieves a user from the database by ID.
func GetUserByID(userID int) (*User, error) {
	// Implement the logic to retrieve the user from the database by ID
	// ...
	return nil, nil
}

// GetUserByUsername retrieves a user from the database by username.
func GetUserByUsername(username string) (*User, error) {
	// Implement the logic to retrieve the user from the database by username
	// ...
	return nil, nil
}
