package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Register routes
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/users/{id}", userHandler).Methods(http.MethodGet)

	// Error handling
	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	r.MethodNotAllowedHandler = http.HandlerFunc(methodNotAllowedHandler)
	r.Use(loggingMiddleware)

	// Start the server
	fmt.Println("Listening on port 8000...")
	http.ListenAndServe(":8000", r)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the homepage!")
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Look up user in database and return user info
	fmt.Fprintf(w, "Retrieving user with ID %s...\n", id)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Page not found", http.StatusNotFound)
}

func methodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}
