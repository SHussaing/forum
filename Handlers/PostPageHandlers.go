package Handlers

import (
	"encoding/json"
	db "forum/Database"
	"html/template"
	"net/http"
	"strconv"
)

func GetPostHandler(w http.ResponseWriter, r *http.Request) {
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

func AddCommentHandler(w http.ResponseWriter, r *http.Request) {
	if !db.HasSessionToken(r) {
		db.DeleteSessionAndRemoveCookie(w, r)
		http.Redirect(w, r, "/Login", http.StatusSeeOther)
		return
	}
	if r.Method == http.MethodPost {
		postIDStr := r.FormValue("post_id")
		postID, err := strconv.Atoi(postIDStr)
		if err != nil {
			handleError(w, http.StatusBadRequest, err)
			return
		}
		content := r.FormValue("content")

		userID, err := db.GetUserIDBySessionToken(r)
		if err != nil {
			handleError(w, http.StatusUnauthorized, err)
			return
		}

		commentID, err := db.AddComment(postID, userID, content)
		if err != nil {
			handleError(w, http.StatusInternalServerError, err)
			return
		}

		username, err := db.GetUsernameByID(userID)
		if err != nil {
			handleError(w, http.StatusInternalServerError, err)
			return
		}

		response := map[string]interface{}{
			"commentID": commentID,
			"username":  username,
			"content":   content,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
