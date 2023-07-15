// backend/db/database.go

package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

// Initialize initializes the database connection.
func Initialize(connectionString string) error {
	var err error
	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

// GetDB returns the database connection.
func GetDB() *sql.DB {
	return db
}
