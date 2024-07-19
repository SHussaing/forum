package handlers

import "net/http"

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "Templates/Register.html")
}
