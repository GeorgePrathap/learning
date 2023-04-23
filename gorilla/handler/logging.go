package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

func main() {
	// Create a new router
	router := http.NewServeMux()

	// Register a handler for the /hello endpoint
	router.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	// Create a new logging handler using os.Stdout as the output destination
	// loggingHandler := handlers.LoggingHandler(os.Stdout, router)

	// Start the server using the logging handler as the main handler
	// http.ListenAndServe(":8080", loggingHandler)

	loggingHandler := handlers.CombinedLoggingHandler(os.Stdout, router)
	http.ListenAndServe(":8080", loggingHandler)

}
