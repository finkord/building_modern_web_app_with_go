package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/finkord/building_modern_web_app_with_go/pkg/config"
	"github.com/finkord/building_modern_web_app_with_go/pkg/handlers"
	"github.com/finkord/building_modern_web_app_with_go/pkg/render"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig

var session *scs.SessionManager

// main is the main function
func main() {

	// Change this to true in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Println("Cannot create template cache")
	}

	app.TemplateCache = tc

	app.UseCache = false

	app.Session = session

	repo := handlers.NewRepo(&app)

	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Printf("Starting server on port %s\n", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: router(&app),
	}

	log.Fatal(srv.ListenAndServe())

}
