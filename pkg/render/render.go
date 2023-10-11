package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func RenderTemplatesOld(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)

	if err != nil {
		fmt.Println("Some error with parsing template", err)
		return
	}
}

var tc = make(map[string]*template.Template)

func RenderTemplates(w http.ResponseWriter, t string) {
	var err error

	_, inMap := tc[t]

	if !inMap {
		//create template and add to cache
		log.Println("Creating new template and add it to cache")
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		//have the tmpl
	}

	err = tc[t].Execute(w, nil)

}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}
	//parse template
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}
	//add template to cache
	tc[t] = tmpl

	return nil
}
