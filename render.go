package render

import (
	"bytes"
	"fmt"
	"lldsdesigns/pkg/config"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

// low case == private only accisible in this code
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	tc := app.TemplateCache

	tc, err := CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal(err)
	}

	// bitbuffer
	buf := new(bytes.Buffer)

	_ = t.Execute(buf, nil)

	_, err = buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing to template to browser", err)
	}
}

var app *config.AppConfig

// Set config for new package
func NewTemplate(a *config.AppConfig) {
	app = a
}

// helps w returning proper data ie todays data outside our templates and fucntions
var functions = template.FuncMap{}

//creates template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {

	// assoctive arrays aka data structure: cache
	// made a map index
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	// controler for loop
	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")

		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")

			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
