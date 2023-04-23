package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	// Create a new router
	r := mux.NewRouter()

	// Define routes and handlers with URL parameters
	r.HandleFunc("/users/{id}", getUserHandler).Methods("GET")
	r.HandleFunc("/users/{id:[0-9]+}", getUserHandler).Methods("GET")
	r.HandleFunc("/users/{id:[0-9]+}/posts", getUserPostsHandler).Methods("GET")
	r.HandleFunc("/users/{id:[0-9]+}/{category}", getUserCategoryHandler).Methods("GET")
	r.HandleFunc("/books/{title}/page/{page:[0-9]+}", getBookPageHandler).Methods("GET")

	// Start the server
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprintf(w, "User ID: %s", id)
}

func getUserPostsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprintf(w, "User ID: %s\nPosts", id)
}

func getUserCategoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	category := vars["category"]
	fmt.Fprintf(w, "User ID: %s\nCategory: %s", id, category)
}

func getBookPageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	pageStr := vars["page"]
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		http.Error(w, "Invalid page number", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Book Title: %s\nPage Number: %d", title, page)
}
