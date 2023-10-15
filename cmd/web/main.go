package main

import (
	"log"
	"net/http"

	"github.com/luciorim/booking/pkg/config"
	"github.com/luciorim/booking/pkg/handlers"
	"github.com/luciorim/booking/pkg/render"
)

const port = ":3000"

func main() {

	var app config.AppConfig

	tmplCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	app.TemplateCache = tmplCache
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	srv := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}
	log.Fatal(srv.ListenAndServe())

}
