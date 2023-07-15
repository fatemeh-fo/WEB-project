// backend/db/database.go

package db

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

// Initialize initializes the database connection and executes the SQL script.
func Initialize(connectionString string, sqlScriptPath string) error {
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

	fmt.Println("Connected to the database")

	// Execute the SQL script
	err = executeSQLScript(sqlScriptPath)
	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println("Database schema created")

	return nil
}

// GetDB returns the database connection.
func GetDB() *sql.DB {
	return db
}

// executeSQLScript reads and executes the SQL script from the given file path.
func executeSQLScript(filePath string) error {
	sqlBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	sqlScript := string(sqlBytes)

	_, err = db.Exec(sqlScript)
	if err != nil {
		return err
	}

	return nil
}
