package respond

import (
	"encoding/json"
	"net/http"

	"goprojstructtest/internal/render"
)

func HTML(w http.ResponseWriter, renderer *render.Renderer, code int, name string, data map[string]any) {
	renderer.HTML(w, code, name, data)
}

func JSON(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}

func Error(w http.ResponseWriter, renderer *render.Renderer, code int, message string) {
	renderer.HTML(w, code, "error.tmpl", map[string]any{
		"Title":   "Error",
		"Message": message,
		"Code":    code,
	})
}
