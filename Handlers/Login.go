package handlers

import (
	db "forum/Database"
	"net/http"
)

// LoginHandler handles the login request
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.ServeFile(w, r, "Templates/Login.html")
		return
	}

	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		password := r.FormValue("password")
		err := db.AuthenticateUser(email, password)
		if err != nil {
			handleError(w, http.StatusBadRequest, err)
			return
		}
		// Redirect to index page
		http.Redirect(w, r, "/", http.StatusSeeOther)

		return
	}
}
