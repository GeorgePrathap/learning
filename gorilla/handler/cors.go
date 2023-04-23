package main

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Define a route that allows cross-origin requests
	router.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	}).Methods("GET")

	// Use the CORS middleware with all options set
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"https://example.com"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowCredentials(),
		handlers.ExposedHeaders([]string{"Content-Length"}),
		handlers.MaxAge(86400),
	)

	http.ListenAndServe(":8080", corsHandler(router))
}

/* The handlers.CORS middleware takes several options, which we pass in as arguments. Here's what each option does:

 	AllowedOrigins([]string{"*"}):
	Allows requests from any origin. You can also specify specific origins by passing in a list of strings.

	AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}):
	Allows requests using any of the specified HTTP methods. You can customize this list as needed.

	AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}):
	Allows requests with any of the specified HTTP headers. Again, you can customize this list as needed.

	AllowCredentials():
	Allows requests with credentials such as cookies or HTTP authentication headers.

	ExposedHeaders([]string{"Content-Length"}):
	Exposes additional HTTP headers in the response that can be accessed by the client.

	MaxAge(86400):
	Sets the amount of time (in seconds) that the CORS preflight response can be cached by the client.

These options allow you to customize how CORS requests are handled by the middleware. For example, you might want to allow requests from only specific origins or restrict the allowed HTTP methods. */
