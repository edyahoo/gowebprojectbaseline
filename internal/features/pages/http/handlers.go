package http

import (
	"net/http"

	"goprojstructtest/internal/render"
)

type Handler struct {
	renderer *render.Renderer
}

func NewHandler(renderer *render.Renderer) *Handler {
	return &Handler{
		renderer: renderer,
	}
}

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	h.renderer.HTML(w, http.StatusOK, "index.tmpl", map[string]any{
		"Title": "Home",
	})
}

func (h *Handler) About(w http.ResponseWriter, r *http.Request) {
	h.renderer.HTML(w, http.StatusOK, "about.tmpl", map[string]any{
		"Title": "About",
	})
}

func (h *Handler) Demo(w http.ResponseWriter, r *http.Request) {
	h.renderer.HTML(w, http.StatusOK, "demo.tmpl", map[string]any{
		"Title": "Demo",
	})
}
