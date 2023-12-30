package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/kunovon/go-kgv-adminstration/pkg/config"
	handlers "github.com/kunovon/go-kgv-adminstration/pkg/handler"
	"github.com/kunovon/go-kgv-adminstration/pkg/render"
)

const portNumber = ":8080"

var app config.AppConfig

//var session *scs.SessionManager

// main is the main function
func main() {

	// change this in production to true
	app.InProduction = false

	// anders als im Beispiel dort session = app.Session
	app.Session = scs.New()
	app.Session.Lifetime = 24 * time.Hour
	app.Session.Cookie.Persist = true
	app.Session.Cookie.SameSite = http.SameSiteLaxMode
	app.Session.Cookie.Secure = app.InProduction

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	//repo := handlers.NewRepo(&app)
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Printf("Staring application on port %s \n", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()

	log.Fatal(err)
}
