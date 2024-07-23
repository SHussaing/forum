package Handlers

import (
	db "forum/Database"
	"html/template"
	"net/http"
	"strconv"
)

var tmpl = template.Must(template.ParseFiles("Templates/CreatePost.html"))

func CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		categories, err := db.GetAllCategories()
		if err != nil {
			handleError(w, http.StatusInternalServerError, err)
			return
		}
		tmpl.Execute(w, struct{ Categories []db.Category }{Categories: categories})
		return
	}

	if r.Method == http.MethodPost {
		title := r.FormValue("title")
		content := r.FormValue("content")
		categories := r.Form["categories"]

		userID, err := db.GetUserIDBySessionToken(r)
		if err != nil {
			handleError(w, http.StatusUnauthorized, err)
			return
		}

		postID, err := db.CreatePost(title, content, userID)
		if err != nil {
			handleError(w, http.StatusInternalServerError, err)
			return
		}

		for _, categoryIDStr := range categories {
			categoryID, err := strconv.Atoi(categoryIDStr)
			if err != nil {
				handleError(w, http.StatusBadRequest, err)
				return
			}
			err = db.AddPostCategory(postID, categoryID)
			if err != nil {
				handleError(w, http.StatusInternalServerError, err)
				return
			}
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
