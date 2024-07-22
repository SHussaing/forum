package Handlers

import (
	"fmt"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		handleError(w, http.StatusNotFound, fmt.Errorf("page not found"))
		return
	}
	http.ServeFile(w, r, "Templates/Index.html")
}
