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

func (h *Handler) LoginPage(w http.ResponseWriter, r *http.Request) {
	h.renderer.HTML(w, http.StatusOK, "login.tmpl", map[string]any{
		"Title":        "Login",
		"ShowError":    false,
		"ErrorMessage": "",
		"Email":        "",
	})
}

func (h *Handler) LoginSubmit(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	if email == "admin@example.com" && password == "password" {
		http.Redirect(w, r, "/demo", http.StatusFound)
		return
	}

	h.renderer.HTML(w, http.StatusOK, "login.tmpl", map[string]any{
		"Title":        "Login",
		"ShowError":    true,
		"ErrorMessage": "Invalid email or password",
		"Email":        email,
	})
}

func (h *Handler) ForgotPasswordPage(w http.ResponseWriter, r *http.Request) {
	h.renderer.HTML(w, http.StatusOK, "forgot-password.tmpl", map[string]any{
		"Title": "Forgot Password",
	})
}

func (h *Handler) ForgotPasswordSubmit(w http.ResponseWriter, r *http.Request) {
	h.renderer.HTML(w, http.StatusOK, "forgot-password.tmpl", map[string]any{
		"Title":   "Forgot Password",
		"Success": true,
		"Message": "If an account exists with this email, you will receive a password reset link shortly.",
	})
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/login", http.StatusFound)
}
