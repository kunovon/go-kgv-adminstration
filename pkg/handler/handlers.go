package handlers

import (
	"net/http"

	"github.com/kunovon/go-kgv-adminstration/pkg/config"
	"github.com/kunovon/go-kgv-adminstration/pkg/models"
	"github.com/kunovon/go-kgv-adminstration/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo create a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandler sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hallo, Markus"

	// send data to the template
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Impressum is the handler for the impressum page
func (m *Repository) Impressum(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "impressum.page.tmpl", &models.TemplateData{})
}
