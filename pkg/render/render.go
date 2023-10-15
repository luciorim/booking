package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/luciorim/booking/pkg/config"
	"github.com/luciorim/booking/pkg/model"
)

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *model.TemplateData) *model.TemplateData {
	return td
}

func RenderTemplates(w http.ResponseWriter, tmpl string, td *model.TemplateData) {
	//create a template cache
	tmplCache := map[string]*template.Template{}
	if app.UseCache {
		tmplCache = app.TemplateCache
	} else {
		tmplCache, _ = CreateTemplateCache()
	}

	//get requested template from cache
	template, ok := tmplCache[tmpl]
	if !ok {
		log.Fatal("Could not get a template")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	_ = template.Execute(buf, td)

	//render the template
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Fatal(err)
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {
	tmplCache := map[string]*template.Template{}

	//get all of the files named *.page.tmpl from ./templates

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return tmplCache, err
	}

	//range through all pages ending with *.page.tmpl

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return tmplCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return tmplCache, err
			}
		}
		tmplCache[name] = ts
	}

	return tmplCache, nil

}
