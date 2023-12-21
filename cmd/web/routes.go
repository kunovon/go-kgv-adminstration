package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/kunovon/go-kgv-adminstration/pkg/config"
	handlers "github.com/kunovon/go-kgv-adminstration/pkg/handler"
)

func routes(app *config.AppConfig) http.Handler {
	// mux := pat.New()

	// mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	// mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	// mux.Get("/impressum", http.HandlerFunc(handlers.Repo.Impressum))

	mux := chi.NewRouter()

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/impressum", handlers.Repo.Impressum)

	return mux
}
