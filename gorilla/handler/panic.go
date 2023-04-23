package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

func main() {
	// Create a new router
	router := http.NewServeMux()

	// Add a handler to the router
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// This will cause a divide by zero error
		var amount []string
		fmt.Printf(amount[10])
		w.Write([]byte("Hello, world!"))
	})

	// Create a new logger
	myLogger := log.New(os.Stderr, "", log.LstdFlags)

	// Define a custom handler function
	// myHandlerFunc := func(w http.ResponseWriter, r *http.Request, err interface{}) {
	// 	fmt.Fprintf(w, "Oops! Something went wrong: %v", err)
	// }

	// Create a new RecoveryHandler middleware with options
	recovery := handlers.RecoveryHandler(
		handlers.PrintRecoveryStack(true),
		handlers.RecoveryLogger(myLogger),
		// handlers.RecoveryHandlerFunc(myHandlerFunc),
	)

	// Wrap the router with the RecoveryHandler middleware
	wrappedRouter := recovery(router)

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", wrappedRouter))
}

// handlers.PrintRecoveryStack(true) prints the stack trace of the panic to the console.
// handlers.RecoveryLogger(myLogger) logs the panic to a custom logger.
// handlers.RecoveryHandlerFunc(myHandlerFunc) calls a custom handler function when a panic occurs.
