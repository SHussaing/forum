package Handlers

import (
	"fmt"
	db "forum/Database"
	"html/template"
	"io"
	"net/http"
	"strconv"
)

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	if !db.HasSessionToken(r) {
		db.DeleteSessionAndRemoveCookie(w, r)
		http.Redirect(w, r, "/Login", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodGet {
		categories, err := db.GetAllCategories()
		if err != nil {
			handleError(w, http.StatusInternalServerError, err)
			return
		}
		tmpl, err := template.ParseFiles("Templates/CreatePost.html")
		if err != nil {
			handleError(w, http.StatusInternalServerError, err)
			return
		}
		tmpl.Execute(w, struct{ Categories []db.Category }{Categories: categories})
		return
	}

	if r.Method == http.MethodPost {
		// Parse form with multipart form to handle file uploads
		err := r.ParseMultipartForm(32 << 20) // Limit to 32MB
		if err != nil {
			handleError(w, http.StatusBadRequest, fmt.Errorf("error parsing form: %v", err))
			return
		}

		title := r.FormValue("title")
		content := r.FormValue("content")
		categories := r.Form["categories"]

		// Retrieve image from form
		file, handler, err := r.FormFile("image")
		var imageData []byte

		if err == nil {
			// Ensure file is closed after processing
			defer file.Close()

			// Validate the image type and size
			allowedTypes := map[string]bool{
				"image/jpeg": true,
				"image/png":  true,
				"image/gif":  true,
			}

			if !allowedTypes[handler.Header.Get("Content-Type")] {
				handleError(w, http.StatusBadRequest, fmt.Errorf("invalid file type"))
				return
			}

			if handler.Size > 20*1024*1024 {
				handleError(w, http.StatusBadRequest, fmt.Errorf("file is too large"))
				return
			}

			// Read image into memory
			imageData, err = io.ReadAll(file)
			if err != nil {
				handleError(w, http.StatusInternalServerError, fmt.Errorf("error reading file: %v", err))
				return
			}
		} else if err != http.ErrMissingFile {
			// Handle any error that isn't due to a missing file
			handleError(w, http.StatusBadRequest, fmt.Errorf("error retrieving file: %v", err))
			return
		}

		userID, err := db.GetUserIDBySessionToken(r)
		if err != nil {
			handleError(w, http.StatusUnauthorized, err)
			return
		}

		// Create post with image data (imageData is nil if no file was uploaded)
		postID, err := db.CreatePost(title, content, userID, imageData)
		if err != nil {
			handleError(w, http.StatusInternalServerError, err)
			return
		}

		// Add categories to post
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
