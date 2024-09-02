package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// main is the entry point for the application
func main() {

	http.HandleFunc("/", Home)

	http.HandleFunc("/about", About)

	fmt.Printf("Starting application on port %s", portNumber)
	http.ListenAndServe(portNumber, nil)
}
