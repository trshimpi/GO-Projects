package main

import (
	"errors"
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

// Divide is the divide page handler
func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100, 10)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err.Error())
		return
	}
	fmt.Fprintf(w, "%d divided by %d is %d", 100, 0, f)
}

func divideValues(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return x / y, nil
}

// main is the entry point for the application
func main() {

	http.HandleFunc("/", Home)

	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Printf("Starting application on port %s", portNumber)
	http.ListenAndServe(portNumber, nil)
}
