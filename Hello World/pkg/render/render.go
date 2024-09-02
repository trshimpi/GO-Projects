package render

import (
	"fmt"
	"html/template"
	"net/http"
)

// renderTemplate renders a template
func RenderTemplateTest(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}
}

var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	var t *template.Template
	var err error

	_, inMap := tc[tmpl]
	if !inMap {
		//need to create the template cache
		fmt.Println("creating template and adding to cache")
		err = createTemplateCache(tmpl)
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		fmt.Println("using cached template")

	}
	t = tc[tmpl]
	err = t.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func createTemplateCache(tmpl string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", tmpl),
		"./templates/base.layout.tmpl",
	}
	t, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}
	tc[tmpl] = t
	return nil
}
