package Handlers

import (
	db "forum/Database"
	"html/template"
	"net/http"
	"strconv"
)

func GetPost(w http.ResponseWriter, r *http.Request) {
	// Get post ID from URL query parameter
	postIDStr := r.URL.Query().Get("id")
	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		handleError(w, http.StatusBadRequest, err)
		return
	}

	// Get the post and its comments
	post, comments, err := db.GetPostAndComments(postID)
	if err != nil {
		handleError(w, http.StatusInternalServerError, err)
		return
	}

	// Render the template
	data := struct {
		Post     db.Post
		Comments []db.Comment
	}{
		Post:     post,
		Comments: comments,
	}
	postTemplate, err := template.ParseFiles("Templates/post.html")
	if err != nil {
		handleError(w, http.StatusInternalServerError, err)
		return
	}
	err = postTemplate.Execute(w, data)
	if err != nil {
		handleError(w, http.StatusInternalServerError, err)
	}
}
