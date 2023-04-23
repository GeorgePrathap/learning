package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Define sub-routers
	productsRouter := r.PathPrefix("/products").Subrouter()
	cartRouter := r.PathPrefix("/cart").Subrouter()
	checkoutRouter := r.PathPrefix("/checkout").Subrouter()

	// Register routes for each sub-router
	productsRouter.HandleFunc("/", listProductsHandler)
	productsRouter.HandleFunc("/{id}", getProductHandler)

	cartRouter.HandleFunc("/", viewCartHandler)
	cartRouter.HandleFunc("/{id}", addToCartHandler)

	checkoutRouter.HandleFunc("/", checkoutHandler)

	// Define a root-level handler
	r.HandleFunc("/", homeHandler)

	// Start the server
	fmt.Println("Listening on port 8000...")
	http.ListenAndServe(":8000", r)
}

func listProductsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Listing products...")
}

func getProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprintf(w, "Getting product %s...", id)
}

func viewCartHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Viewing cart...")
}

func addToCartHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprintf(w, "Adding product %s to cart...", id)
}

func checkoutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Checking out...")
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the homepage!")
}
