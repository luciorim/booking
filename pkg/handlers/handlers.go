package handlers

import (
	"net/http"

	"github.com/luciorim/booking/pkg/config"
	"github.com/luciorim/booking/pkg/model"
	"github.com/luciorim/booking/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// Sets repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, "home.page.tmpl", &model.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "privet mir"

	//send data to to the template
	render.RenderTemplates(w, "about.page.tmpl", &model.TemplateData{
		StringMap: stringMap,
	})
}
