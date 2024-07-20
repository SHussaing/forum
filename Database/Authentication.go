package Database

import (
	"errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

// InsertUser inserts a new user into the User table
func InsertUser(email, username, password string) error {
	// Check if email or username already exists
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM User WHERE email = ? OR username = ?)`
	err := db.QueryRow(query, email, username).Scan(&exists)
	if err != nil {
		return fmt.Errorf("failed to check if user exists: %v", err)
	}
	if exists {
		return errors.New("email or username already exists")
	}

	// Encrypt the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to encrypt password: %v", err)
	}

	// Insert the new user
	insertUserSQL := `INSERT INTO User (email, username, password) VALUES (?, ?, ?)`
	_, err = db.Exec(insertUserSQL, email, username, hashedPassword)
	if err != nil {
		return fmt.Errorf("failed to insert user: %v", err)
	}

	return nil
}
