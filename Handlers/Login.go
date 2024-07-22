package Handlers

import (
	db "forum/Database"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func generateSessionToken() (string, error) {
	return uuid.New().String(), nil
}

// LoginHandler handles the login request
func LoginHandler(w http.ResponseWriter, r *http.Request) {
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
		token, err := generateSessionToken()
		if err != nil {
			handleError(w, http.StatusInternalServerError, err)
			return
		}

		// Set session expiration (e.g., 24 hours)
		expiresAt := time.Now().Add(24 * time.Hour)

		// Store the session in the database and set the cookie
		if err := db.CreateSessionAndSetCookie(w, userID, token, expiresAt); err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// Redirect to the index page
		http.Redirect(w, r, "/", http.StatusSeeOther)

	}

}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	_, err = db.Db.Exec("DELETE FROM Sessions WHERE token = ?", cookie.Value)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Clear the cookie
	http.SetCookie(w, &http.Cookie{
		Name:   "session_token",
		Value:  "",
		MaxAge: -1,
		Path:   "/",
	})

	http.Redirect(w, r, "/login", http.StatusFound)
}
