package main

import (
	"fmt"
	handlers "forum/Handlers"
	"log"
	"net/http"
)

func main() {
	// Define the directory to serve static files from
	http.Handle("/Static/", http.StripPrefix("/Static/", http.FileServer(http.Dir("Static"))))

	// Define the routes
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/Login", handlers.LoginHandler)
	http.HandleFunc("/Logout", handlers.LogoutHandler)
	http.HandleFunc("/Register", handlers.RegisterHandler)
	http.HandleFunc("/CreatePost", handlers.CreatePostHandler)
	http.HandleFunc("/Post", handlers.GetPostHandler)
	http.HandleFunc("/AddComment", handlers.AddCommentHandler)
	http.HandleFunc("/LikeDislike", handlers.LikeDislikeHandler)
	http.HandleFunc("/Filter", handlers.FilterHandler)
	http.HandleFunc("/ServeImage", handlers.ServeImageHandler)
	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}