package Handlers

import (
	"net/http"

	"github.com/google/uuid"
)

// Function to generate a session token
func GenerateSessionToken() (string, error) {
	return uuid.New().String(), nil
}

// Function to check if the session_token cookie exists
func HasSessionToken(r *http.Request) bool {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		return false
	}
	return cookie != nil
}
