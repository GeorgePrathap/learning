package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/gorilla/mux"
)

func main() {
	// Create a new router
	r := mux.NewRouter()

	// Define routes and handlers with advanced URL patterns and matchers
	r.HandleFunc("/users/{id:[0-9]+}[/{category}]", getUserCategoryHandler).Methods("GET")
	r.HandleFunc("/products/{uuid}", getProductHandler).Methods("GET").MatcherFunc(uuidMatcher)

	// Create a subrouter for blog posts
	blogRouter := r.PathPrefix("/blog").Subrouter()
	blogRouter.HandleFunc("/", getAllPostsHandler).Methods("GET")
	blogRouter.HandleFunc("/{slug}", getPostHandler).Methods("GET")

	// Start the server
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func getUserCategoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	category, ok := vars["category"]
	if !ok {
		category = "all"
	}
	fmt.Fprintf(w, "User ID: %s\nCategory: %s", id, category)
}

func getProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["uuid"]
	fmt.Fprintf(w, "Product UUID: %s", uuid)
}

func uuidMatcher(r *http.Request, rm *mux.RouteMatch) bool {
	uuidRegexp := `[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}`
	regex := fmt.Sprintf("^%s$", uuidRegexp)
	if matched, _ := regexp.MatchString(regex, rm.Vars["uuid"]); matched {
		return true
	}
	return false
}

func getAllPostsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "All posts")
}

func getPostHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug := vars["slug"]
	fmt.Fprintf(w, "Post slug: %s", slug)
}
