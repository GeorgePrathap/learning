package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Middleware function for authentication
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if user is authenticated
		if isAuthenticated(r) {
			// Call the next handler if authenticated
			next.ServeHTTP(w, r)
		} else {
			// Redirect to login page if not authenticated
			http.Redirect(w, r, "/login", http.StatusFound)
		}
	})
}

// Middleware function for logging
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}

func main() {
	r := mux.NewRouter()

	// Register routes
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/profile", profileHandler)

	// Apply middleware to specific routes
	r.HandleFunc("/admin", adminHandler).Methods(http.MethodGet).MiddlewareFunc(authMiddleware, loggingMiddleware)

	// Apply middleware to all routes
	r.Use(loggingMiddleware)

	// Start the server
	fmt.Println("Listening on port 8000...")
	http.ListenAndServe(":8000", r)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the homepage!")
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Viewing your profile...")
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Viewing the admin panel...")
}

// Helper function to check if the user is authenticated
func isAuthenticated(r *http.Request) bool {
	// Check the user's session or cookie for authentication
	return true
}
