package handlers

import (
	"net/http"

	"github.com/fabiobap/go-tibia-calc/internal/config"
	"github.com/fabiobap/go-tibia-calc/internal/models"
	"github.com/fabiobap/go-tibia-calc/internal/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(ac *config.AppConfig) *Repository {
	return &Repository{
		App: ac,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "home.page.tmpl", &models.TemplateData{})
}
