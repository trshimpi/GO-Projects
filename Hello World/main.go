package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum, err := addValues(2, 3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w, "The sum of %d and %d is %d", 2, 3, sum)
}

// addValues is a helper function to add two values
func addValues(x, y int) (int, error) {
	return x + y, nil
}

// main is the entry point for the application
func main() {

	http.HandleFunc("/", Home)

	http.HandleFunc("/about", About)

	fmt.Printf("Starting application on port %s", portNumber)
	http.ListenAndServe(portNumber, nil)
}
