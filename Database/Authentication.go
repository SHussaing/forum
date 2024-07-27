package Database

import (
	"database/sql"
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
	err := Db.QueryRow(query, email, username).Scan(&exists)
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
	_, err = Db.Exec(insertUserSQL, email, username, hashedPassword)
	if err != nil {
		return fmt.Errorf("failed to insert user: %v", err)
	}

	return nil
}

// validateUserCredentials validates the user credentials and returns the user ID
func ValidateUserCredentials(email, password string) (int, error) {
	var userID int
	var storedPassword string
	err := Db.QueryRow("SELECT user_ID, password FROM User WHERE email = ?", email).Scan(&userID, &storedPassword)
	if err != nil {
		return 0, errors.New("invalid email or password")
	}

	// Compare the stored password with the provided password
	if err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password)); err != nil {
		return 0, err
	}

	return userID, nil
}

// Function to create a session and set a cookie
func CreateSessionAndSetCookie(w http.ResponseWriter, userID int, token string, expiresAt time.Time) error {
	// Delete any existing sessions for this user
	_, err := Db.Exec("DELETE FROM Session WHERE user_ID = ?", userID)
	if err != nil {
		return err
	}

	// Store the session in the database
	_, err = Db.Exec("INSERT INTO Session (user_ID, token, expires_at) VALUES (?, ?, ?)", userID, token, expiresAt)
	if err != nil {
		return err
	}

	// Set the session token as a cookie, accessible site-wide
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   token,
		Expires: expiresAt,
		Path:    "/", // Make the cookie accessible across the entire site
	})

	return nil
}

// Function to delete the session and remove the cookie
func DeleteSessionAndRemoveCookie(w http.ResponseWriter, r *http.Request) error {
	// Get the session_token cookie
	cookie, err := r.Cookie("session_token")
	if err != nil {
		// If the cookie does not exist, nothing to do
		return nil
	}

	// Delete the session from the database
	_, err = Db.Exec("DELETE FROM Session WHERE token = ?", cookie.Value)
	if err != nil {
		return err
	}

	// Remove the cookie by setting it with an expired date
	http.SetCookie(w, &http.Cookie{
		Name:    cookie.Name,
		Value:   "",
		Expires: time.Unix(0, 0),
		Path:    "/",
	})

	return nil
}

func GetUserIDBySessionToken(r *http.Request) (int, error) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		return 0, errors.New("session token not found")
	}

	var userID int
	err = Db.QueryRow("SELECT user_ID FROM Session WHERE token = ?", cookie.Value).Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, errors.New("session token not found")
		}
		return 0, err
	}
	return userID, nil
}

// Function to check if the session_token cookie exists and is valid in the database
func HasSessionToken(r *http.Request) bool {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		return false
	}

	// If the cookie is found, check if it exists in the database
	var userID int
	err = Db.QueryRow("SELECT user_ID FROM Session WHERE token = ?", cookie.Value).Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return false // Token not found in database
		}
		// Handle other potential errors
		return false
	}
	return true
}
