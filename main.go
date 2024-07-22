package main

import (
	"fmt"
	"forum/handlers"
	"log"
	"net/http"
)

func main() {
	// Define the directory to serve static files from
	http.Handle("/Static/", http.StripPrefix("/Static/", http.FileServer(http.Dir("Static"))))

	// Define the routes
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/Login", handlers.LoginHandler)
	http.HandleFunc("/Register", handlers.RegisterHandler)

	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
