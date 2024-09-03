package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/trshimpi/GO-projects/pkg/config"
)

// // renderTemplate renders a template without caching
// func RenderTemplateTest(w http.ResponseWriter, tmpl string) {
// 	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
// 	err := parsedTemplate.Execute(w, nil)
// 	if err != nil {
// 		fmt.Println("Error parsing template:", err)
// 		return
// 	}
// }

// var tc = make(map[string]*template.Template)

// // Simple approach to caching templates and adding them to the template cache
// func RenderTemplate(w http.ResponseWriter, tmpl string) {
// 	var t *template.Template
// 	var err error

// 	_, inMap := tc[tmpl]
// 	if !inMap {
// 		//need to create the template cache
// 		fmt.Println("creating template and adding to cache")
// 		err = createTemplateCache(tmpl)
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}
// 	} else {
// 		fmt.Println("using cached template")

// 	}
// 	t = tc[tmpl]
// 	err = t.Execute(w, nil)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// }

//	func createTemplateCache(tmpl string) error {
//		templates := []string{
//			fmt.Sprintf("./templates/%s", tmpl),
//			"./templates/base.layout.tmpl",
//		}
//		t, err := template.ParseFiles(templates...)
//		if err != nil {
//			return err
//		}
//		tc[tmpl] = t
//		return nil
//	}

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, tmpl string) {

	//get templatecache from Appconfig
	tc := app.TemplateCache

	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template fron template cache")
	}

	buf := new(bytes.Buffer)
	err := t.Execute(buf, nil)

	if err != nil {
		log.Println(err)
	}

	//render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all of the files named *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return myCache, err
	}

	//range through all files ending with *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}
	return myCache, nil
}
