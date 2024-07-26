package Handlers

import (
	"fmt"
	"html/template"
	"net/http"
	db "forum/Database"
)

func FilterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		handleError(w, http.StatusMethodNotAllowed, fmt.Errorf("method not allowed"))
		return
	}

	r.ParseForm()

	userID, err := db.GetUserIDBySessionToken(r)
	isLoggedIn := err == nil

	filters := r.Form["filter"]
	categoryIDs := r.Form["category"]

	posts, err := db.GetFilteredPosts(userID, filters, categoryIDs, isLoggedIn)
	if err != nil {
		handleError(w, http.StatusInternalServerError, err)
		return
	}

	categories, err := db.GetAllCategories()
	if err != nil {
		handleError(w, http.StatusInternalServerError, err)
		return
	}

	data := IndexPageData{
		Posts:      posts,
		Categories: categories,
	}

	tmpl, err := template.ParseFiles("Templates/index.html")
	if err != nil {
		handleError(w, http.StatusInternalServerError, err)
		return
	}
	tmpl.Execute(w, data)
}
