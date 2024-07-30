package Handlers

import (
	"fmt"
	db "forum/Database"
	"html/template"
	"net/http"
)

type IndexPageData struct {
	Posts      []db.Post
	Categories []db.Category
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		handleError(w, http.StatusNotFound, fmt.Errorf("page not found"))
		return
	}

	posts, err := db.GetAllPosts()
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

	tmpl, err := template.ParseFiles("Templates/Index.html")
	if err != nil {
		handleError(w, http.StatusInternalServerError, err)
		return
	}
	tmpl.Execute(w, data)
}
