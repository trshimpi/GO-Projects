package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/trshimpi/GO-projects/pkg/config"
	"github.com/trshimpi/GO-projects/pkg/handlers"
	"github.com/trshimpi/GO-projects/pkg/render"
)

const portNumber = ":8080"

// main is the entry point for the application
func main() {

	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannont create template cache")
	}
	app.TemplateCache = tc

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Home)

	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port %s", portNumber)
	http.ListenAndServe(portNumber, nil)
}
