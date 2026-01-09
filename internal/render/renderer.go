package render

import (
	"html/template"
	"net/http"
	"path/filepath"
)

type Renderer struct {
	templates *template.Template
}

func New() (*Renderer, error) {
	tmpl := template.New("")

	// Load layout templates
	layoutPattern := filepath.Join("web", "templates", "layouts", "*.tmpl")
	layouts, err := filepath.Glob(layoutPattern)
	if err != nil {
		return nil, err
	}
	if len(layouts) > 0 {
		if _, err := tmpl.ParseFiles(layouts...); err != nil {
			return nil, err
		}
	}

	// Load page templates
	pagePattern := filepath.Join("web", "templates", "pages", "*.tmpl")
	pages, err := filepath.Glob(pagePattern)
	if err != nil {
		return nil, err
	}
	if len(pages) > 0 {
		if _, err := tmpl.ParseFiles(pages...); err != nil {
			return nil, err
		}
	}

	// Load partial templates (optional - may be empty)
	partialPattern := filepath.Join("web", "templates", "partials", "*.tmpl")
	partials, err := filepath.Glob(partialPattern)
	if err != nil {
		return nil, err
	}
	if len(partials) > 0 {
		if _, err := tmpl.ParseFiles(partials...); err != nil {
			return nil, err
		}
	}

	return &Renderer{templates: tmpl}, nil
}

func (r *Renderer) Template() *template.Template {
	return r.templates
}

func (r *Renderer) HTML(w http.ResponseWriter, status int, name string, data any) error {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)
	return r.templates.ExecuteTemplate(w, name, data)
}
