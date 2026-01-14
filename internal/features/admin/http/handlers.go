package http

import (
	"log/slog"
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

func (h *Handler) Dashboard(w http.ResponseWriter, r *http.Request) {
	slog.Info("Admin Dashboard handler called", "path", r.URL.Path)
	err := h.renderer.HTML(w, http.StatusOK, "dashboard.tmpl", map[string]any{
		"Title": "Admin Dashboard - EFTG",
	})
	if err != nil {
		slog.Error("Failed to render admin dashboard", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (h *Handler) TestHTML(w http.ResponseWriter, r *http.Request) {
	slog.Info("Serving test.html", "path", r.URL.Path)
	http.ServeFile(w, r, "web/templates/pages/admin/test.html")
}
