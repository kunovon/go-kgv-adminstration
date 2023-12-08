package handlers

import (
	"net/http"

	"github.com/kunovon/go-kgv-adminstration/pkg/render"
)

// Home is the handler for the home page
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

// About is the handler for the about page
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}

// Impressum is the handler for the impressum page
func Impressum(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "impressum.page.tmpl")
}
