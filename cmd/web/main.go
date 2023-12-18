package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kunovon/go-kgv-adminstration/pkg/config"
	handlers "github.com/kunovon/go-kgv-adminstration/pkg/handler"
	"github.com/kunovon/go-kgv-adminstration/pkg/render"
)

const portNumber = ":8080"

// main is the main function
func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache!")
	}

	app.TemplateCache = tc

	app.UseCache = true

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
	http.HandleFunc("/impressum", handlers.Repo.Impressum)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
