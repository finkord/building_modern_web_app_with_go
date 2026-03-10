package main

import (
	"go-monitoring/pkg/config"
	"go-monitoring/pkg/handlers"
	"net/http"

	"github.com/bmizerany/pat"
)

func router(app *config.AppConfig) http.Handler {

	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}
