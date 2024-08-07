package Handlers

import (
	"fmt"
	db "forum/Database"
	"net/http"
)

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

		// Create a session and set the cookie
		if err := db.CreateSessionAndSetCookie(w, userID); err != nil {
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
