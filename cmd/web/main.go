package main

import (
	"fmt"
	"go-monitoring/pkg/config"
	"go-monitoring/pkg/handlers"
	"go-monitoring/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

// main is the main function
func main() {

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Println("Cannot create template cache")
	}

	app.TemplateCache = tc

	app.UseCache = false

	repo := handlers.NewRepo(&app)

	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Starting server on port %s\n", portNumber)
	_ = http.ListenAndServe(portNumber, nil)

}
