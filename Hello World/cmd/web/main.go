package main

import (
	"fmt"
	"net/http"

	"github.com/trshimpi/GO-projects/pkg/handlers"
)

const portNumber = ":8080"

// main is the entry point for the application
func main() {

	http.HandleFunc("/", handlers.Home)

	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port %s", portNumber)
	http.ListenAndServe(portNumber, nil)
}
