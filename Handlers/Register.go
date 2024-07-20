package handlers

import (
	db "forum/Database"
	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.ServeFile(w, r, "Templates/Register.html")
		return
	}

	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		username := r.FormValue("username")
		password := r.FormValue("password")
		err := db.InsertUser(email, username, password)
		if err != nil {
			handleError(w, http.StatusBadRequest, err)
			return
		}

		// Redirect to index page
		http.Redirect(w, r, "/", http.StatusSeeOther)

		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}
