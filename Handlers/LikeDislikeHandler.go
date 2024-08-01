package Handlers

import (
	"encoding/json"
	"fmt"
	db "forum/Database"
	"net/http"
	"strconv"
)

func LikeDislikeHandler(w http.ResponseWriter, r *http.Request) {
	if !db.HasSessionToken(r) {
		db.DeleteSessionAndRemoveCookie(w, r)
		http.Redirect(w, r, "/Login", http.StatusSeeOther)
		return
	}
	if r.Method == http.MethodPost {
		r.ParseForm()
		userID, err := db.GetUserIDBySessionToken(r)
		if err != nil {
			handleError(w, http.StatusUnauthorized, err)
			return
		}

		itemType := r.URL.Query().Get("type")
		itemIDStr := r.URL.Query().Get("id")
		action := r.URL.Query().Get("action")

		itemID, err := strconv.Atoi(itemIDStr)
		if err != nil {
			handleError(w, http.StatusBadRequest, err)
			return
		}

		var likes, dislikes int
		if itemType == "post" {
			if action == "like" {
				err = db.LikePost(userID, itemID)
			} else {
				err = db.DislikePost(userID, itemID)
			}
			if err == nil {
				likes, dislikes, err = db.GetPostLikes(itemID)
			}
		} else if itemType == "comment" {
			if action == "like" {
				err = db.LikeComment(userID, itemID)
			} else {
				err = db.DislikeComment(userID, itemID)
			}
			if err == nil {
				likes, dislikes, err = db.GetCommentLikes(itemID)
			}
		}

		if err != nil {
			handleError(w, http.StatusInternalServerError, err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response := map[string]int{"likes": likes, "dislikes": dislikes}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		handleError(w, http.StatusMethodNotAllowed, fmt.Errorf("method not allowed"))
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
