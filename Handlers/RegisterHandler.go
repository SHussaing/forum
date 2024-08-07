package Handlers

import (
	"fmt"
	db "forum/Database"
	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if db.HasSessionToken(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodGet {
		http.ServeFile(w, r, "Templates/Register.html")
		return
	}

	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Insert user into the database
		userID, err := db.InsertUser(email, username, password)
		if err != nil {
			handleError(w, http.StatusBadRequest, err)
			return
		}

		// Create a session and set the cookie
		if err := db.CreateSessionAndSetCookie(w, int(userID)); err != nil {
			handleError(w, http.StatusInternalServerError, err)
			return
		}

		// Redirect to the index page
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	handleError(w, http.StatusMethodNotAllowed, fmt.Errorf("method not allowed"))
}
