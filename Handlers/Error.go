package Handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

var errorPageTemplate *template.Template

type ErrorPageData struct {
	Title      string // Title of the HTML page
	StatusCode int    // HTTP status code
	Message    string // Error message
}

// Function to create ErrorPageData for the error
func createErrorPageData(statusCode int, err error) ErrorPageData {
	return ErrorPageData{
		Title:      fmt.Sprintf("Error %d", statusCode),
		StatusCode: statusCode,
		Message:    err.Error(),
	}
}

// Function to handle errors and return appropriate HTTP response
func handleError(w http.ResponseWriter, statusCode int, err error) {
	errorData := createErrorPageData(statusCode, err)
	w.WriteHeader(statusCode)
	if err := errorPageTemplate.Execute(w, errorData); err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

func init() {
	var err error
	// Parse the template file
	errorPageTemplate, err = template.ParseFiles(filepath.Join("Templates", "Error.html"))
	if err != nil {
		panic(fmt.Sprintf("Failed to parse template: %v", err))
	}
}
