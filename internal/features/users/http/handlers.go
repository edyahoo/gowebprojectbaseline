package http

import "net/http"

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
}
