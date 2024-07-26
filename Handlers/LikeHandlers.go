package Handlers

import (
	"fmt"
	db "forum/Database"
	"net/http"
	"strconv"
)

func LikePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		postIDStr := r.FormValue("post_id")
		postID, err := strconv.Atoi(postIDStr)
		if err != nil {
			handleError(w, http.StatusBadRequest, err)
			return
		}
		// Handle the like functionality here
		userID, err := db.GetUserIDBySessionToken(r)
		if err != nil {
			handleError(w, http.StatusUnauthorized, err)
			return
		}
		err = db.LikePost(userID, postID)
		if err != nil {
			handleError(w, http.StatusInternalServerError, err)
			return
		}
		http.Redirect(w, r, "/Post?id="+postIDStr, http.StatusSeeOther)
	} else {
		handleError(w, http.StatusMethodNotAllowed, fmt.Errorf("method not allowed"))
	}
}

func DislikePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		postIDStr := r.FormValue("post_id")
		postID, err := strconv.Atoi(postIDStr)
		if err != nil {
			handleError(w, http.StatusBadRequest, err)
			return
		}
		// Handle the dislike functionality here
		userID, err := db.GetUserIDBySessionToken(r)
		if err != nil {
			handleError(w, http.StatusUnauthorized, err)
			return
		}
		err = db.DislikePost(userID, postID)
		if err != nil {
			handleError(w, http.StatusInternalServerError, err)
			return
		}
		http.Redirect(w, r, "/Post?id="+postIDStr, http.StatusSeeOther)
	} else {
		handleError(w, http.StatusMethodNotAllowed, fmt.Errorf("method not allowed"))

	}
}

func LikeCommentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()                       
		userID, err := db.GetUserIDBySessionToken(r)
		if err != nil {
			handleError(w, http.StatusUnauthorized, err)
			return
		}
		commentIDStr := r.FormValue("comment_id")
		commentID, err := strconv.Atoi(commentIDStr)
		if err != nil {
			handleError(w, http.StatusBadRequest, err)
			return
		}
		err = db.LikeComment(userID, commentID)
		if err != nil {
			handleError(w, http.StatusInternalServerError, err)
			return
		}
		postIDStr := r.FormValue("post_id")                
		http.Redirect(w, r, "/Post?id="+postIDStr, http.StatusSeeOther)
	} else {
		handleError(w, http.StatusMethodNotAllowed, fmt.Errorf("method not allowed"))
	}
}

func DislikeCommentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()                       
		userID, err := db.GetUserIDBySessionToken(r)
		if err != nil {
			handleError(w, http.StatusUnauthorized, err)
			return
		}
		commentIDStr := r.FormValue("comment_id")
		commentID, err := strconv.Atoi(commentIDStr)
		if err != nil {
			handleError(w, http.StatusBadRequest, err)
			return
		}
		err = db.DislikeComment(userID, commentID)
		if err != nil {
			handleError(w, http.StatusInternalServerError, err)
			return
		}
		postIDStr := r.FormValue("post_id")
		http.Redirect(w, r, "/Post?id="+postIDStr, http.StatusSeeOther)
	} else {
		handleError(w, http.StatusMethodNotAllowed, fmt.Errorf("method not allowed"))
	}
}
