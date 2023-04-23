package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Create a new router
	r := mux.NewRouter()

	// Define routes and handlers
	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/users", usersHandler).Methods("GET")
	r.HandleFunc("/users", createUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", getUserHandler).Methods("GET")
	r.HandleFunc("/users/{id}", updateUserHandler).Methods("PUT")
	r.HandleFunc("/users/{id}", deleteUserHandler).Methods("DELETE")

	// Start the server
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to my website!")
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "List of all users")
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Create a new user")
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprintf(w, "Get user with ID %s\n", id)
}

func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprintf(w, "Update user with ID %s\n", id)
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprintf(w, "Delete user with ID %s\n", id)
}
