package handlers

import (
	db "forum/database"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func generateSessionToken() (string, error) {
	return uuid.New().String(), nil
}

// LoginHandler handles the login request
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	// Validate user credentials
	userID, err := db.ValidateUserCredentials(email, password)
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// Generate a session token
	token, err := generateSessionToken()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Set session expiration (e.g., 24 hours)
	expiresAt := time.Now().Add(24 * time.Hour)

	// Store the session in the database and set the cookie
	if err := db.CreateSessionAndSetCookie(w, userID, token, expiresAt); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Redirect to a protected page
	http.Redirect(w, r, "/protected", http.StatusFound)
}
