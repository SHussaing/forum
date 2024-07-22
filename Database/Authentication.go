package database

import (
	"errors"
	"fmt"
	"net/http"
	"time"

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

// validateUserCredentials validates the user credentials and returns the user ID
func ValidateUserCredentials(email, password string) (int, error) {
	var userID int
	var storedPassword string
	err := db.QueryRow("SELECT user_ID, password FROM User WHERE email = ?", email).Scan(&userID, &storedPassword)
	if err != nil {
		return 0, err
	}

	// Compare the stored password with the provided password
	if err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password)); err != nil {
		return 0, err
	}

	return userID, nil
}

// createSessionAndSetCookie stores the session in the database and sets the session token as a cookie
func CreateSessionAndSetCookie(w http.ResponseWriter, userID int, token string, expiresAt time.Time) error {
	// Delete any existing sessions for this user
	_, err := db.Exec("DELETE FROM Sessions WHERE user_ID = ?", userID)
	if err != nil {
		return err
	}

	// Store the session in the database
	_, err = db.Exec("INSERT INTO Sessions (user_ID, token, expires_at) VALUES (?, ?, ?)", userID, token, expiresAt)
	if err != nil {
		return err
	}

	// Set the session token as a cookie
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   token,
		Expires: expiresAt,
		Path:    "/",
	})

	return nil
}
