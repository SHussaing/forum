package Handlers

import (
	"fmt"
	db "forum/Database"
	"net/http"
	"time"

	"github.com/google/uuid"
)

// Function to generate a session token
func GenerateSessionToken() (string, error) {
	return uuid.New().String(), nil
}

// LoginHandler handles the login request
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if db.HasSessionToken(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if r.Method == http.MethodGet {
		http.ServeFile(w, r, "Templates/Login.html")
		return
	}

	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		password := r.FormValue("password")

		// Validate user credentials
		userID, err := db.ValidateUserCredentials(email, password)
		if err != nil {
			handleError(w, http.StatusUnauthorized, err)
			return
		}

		// Generate a session token
		token, err := GenerateSessionToken()
		if err != nil {
			handleError(w, http.StatusInternalServerError, err)
			return
		}

		// Set session expiration (6 hours)
		expiresAt := time.Now().Add(6 * time.Hour)

		// Store the session in the database and set the cookie
		if err := db.CreateSessionAndSetCookie(w, userID, token, expiresAt); err != nil {
			handleError(w, http.StatusInternalServerError, err)
			return
		}

		// Redirect to the index page
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	handleError(w, http.StatusMethodNotAllowed, fmt.Errorf("method not allowed"))
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	err := db.DeleteSessionAndRemoveCookie(w, r)
	if err != nil {
		handleError(w, http.StatusInternalServerError, err)
		return
	}

	// Redirect to the index page
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
