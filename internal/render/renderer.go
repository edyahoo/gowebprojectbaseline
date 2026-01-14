package render

import (
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"path/filepath"

	"goprojstructtest/internal/platform/config"
)

type Renderer struct {
	templates *template.Template
	config    *config.Config
}

func New(cfg *config.Config) (*Renderer, error) {
	tmpl, err := loadTemplates()
	if err != nil {
		return nil, err
	}

	return &Renderer{
		templates: tmpl,
		config:    cfg,
	}, nil
}

func loadTemplates() (*template.Template, error) {
	// Collect all template files
	var allFiles []string

	// Load layout templates
	layoutPattern := filepath.Join("web", "templates", "layouts", "*.tmpl")
	layouts, err := filepath.Glob(layoutPattern)
	if err != nil {
		return nil, err
	}
	allFiles = append(allFiles, layouts...)

	// Load page templates
	pagePattern := filepath.Join("web", "templates", "pages", "*.tmpl")
	pages, err := filepath.Glob(pagePattern)
	if err != nil {
		return nil, err
	}
	allFiles = append(allFiles, pages...)

	// Load .html page templates
	pageHTMLPattern := filepath.Join("web", "templates", "pages", "*.html")
	pagesHTML, err := filepath.Glob(pageHTMLPattern)
	if err != nil {
		return nil, err
	}
	allFiles = append(allFiles, pagesHTML...)

	// Load page templates from subdirectories
	pageSubdirPattern := filepath.Join("web", "templates", "pages", "*", "*.tmpl")
	pagesSubdir, err := filepath.Glob(pageSubdirPattern)
	if err != nil {
		return nil, err
	}
	allFiles = append(allFiles, pagesSubdir...)

	// Load .html page templates from subdirectories
	pageSubdirHTMLPattern := filepath.Join("web", "templates", "pages", "*", "*.html")
	pagesSubdirHTML, err := filepath.Glob(pageSubdirHTMLPattern)
	if err != nil {
		return nil, err
	}
	allFiles = append(allFiles, pagesSubdirHTML...)

	// Load partial templates (optional - may be empty)
	partialPattern := filepath.Join("web", "templates", "partials", "*.tmpl")
	partials, err := filepath.Glob(partialPattern)
	if err != nil {
		return nil, err
	}
	allFiles = append(allFiles, partials...)

	// Parse all templates together
	tmpl, err := template.ParseFiles(allFiles...)
	if err != nil {
		return nil, err
	}

	// Debug: Log all loaded template names
	for _, t := range tmpl.Templates() {
		// Only log in debug builds or controlled by env var if desired
		_ = t.Name() // Available for debugging
	}

	return tmpl, nil
}

func (r *Renderer) Template() *template.Template {
	return r.templates
}

func (r *Renderer) HTML(w http.ResponseWriter, status int, name string, data any) error {
	slog.Info("Rendering template", "name", name, "env", r.config.Env)

	// In development mode, reload templates on each request for hot-reload
	if r.config.IsDevelopment() {
		slog.Debug("Development mode: reloading templates")
		tmpl, err := loadTemplates()
		if err != nil {
			slog.Error("Failed to reload templates", "error", err)
			return fmt.Errorf("failed to reload templates: %w", err)
		}
		r.templates = tmpl
	}

	// Debug: Check if template exists
	tmpl := r.templates.Lookup(name)
	if tmpl == nil {
		slog.Error("Template not found", "name", name)
		// List available templates
		slog.Info("Available templates:")
		for _, t := range r.templates.Templates() {
			slog.Info("  - " + t.Name())
		}
		return fmt.Errorf("template not found: %s", name)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)
	err := r.templates.ExecuteTemplate(w, name, data)
	if err != nil {
		slog.Error("Template execution failed", "name", name, "error", err)
	}
	return err
}
