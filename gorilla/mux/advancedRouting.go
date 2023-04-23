package main

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Route with regular expression match
	r.HandleFunc("/user/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		fmt.Fprintf(w, "User ID: %s", id)
	})

	// Route with path prefix match
	r.PathPrefix("/files/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Files endpoint")
	})

	// Route with custom matcher
	slugMatcher := func(r *http.Request, rm *mux.RouteMatch) bool {
		slug := rm.Vars["slug"]
		// Check if slug contains only alphabets and hyphens
		return regexp.MustCompile(`^[a-zA-Z\-]+$`).MatchString(slug)
	}

	r.HandleFunc("/{slug}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Slug: %s", mux.Vars(r)["slug"])
	}).Matchers(slugMatcher).Methods(http.MethodGet)

	// Route with multiple matchers
	r.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Search endpoint")
	}).Methods(http.MethodGet).Queries("q", "{query}")

	// Start the server
	fmt.Println("Listening on port 8000...")
	http.ListenAndServe(":8000", r)
}

/*
Regular expressions:
The route /user/{id:[0-9]+} uses a regular expression to match any numeric value for the id parameter. The regular expression [0-9]+ matches one or more digits. This allows us to create a more specific route that only matches numeric values for the id parameter.
*/

/*
Path prefix match:
The route /files/ uses the PathPrefix method to match any path that starts with /files/. This is useful when you want to group together routes that have a common prefix.
*/

/*
Custom matcher:
The route /{slug} uses a custom matcher function to match any path that contains only alphabets and hyphens for the slug parameter. The Matchers method is used to apply the custom matcher function to this route.
*/

/*
Multiple matchers:
The route /search uses the Methods and Queries methods to match GET requests with a query parameter q. This allows us to create a more specific route that only matches GET requests with the q parameter.
*/
