package http

import (
	"net/http"
	"time"

	"goprojstructtest/internal/domain"
	"goprojstructtest/internal/platform/config"
	"goprojstructtest/internal/platform/session"
	"goprojstructtest/internal/render"
)

type Handler struct {
	renderer *render.Renderer
	store    session.SessionStore
	config   *config.Config
}

func NewHandler(renderer *render.Renderer, store session.SessionStore, cfg *config.Config) *Handler {
	return &Handler{
		renderer: renderer,
		store:    store,
		config:   cfg,
	}
}

func (h *Handler) LoginPage(w http.ResponseWriter, r *http.Request) {
	h.renderer.HTML(w, http.StatusOK, "login.tmpl", map[string]any{
		"Title":        "Login",
		"ShowError":    false,
		"ErrorMessage": "",
		"Email":        "admin@example.com",
		"Password":     "password",
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
		// Create session
		sessionID, err := h.store.Create(
			domain.UserID(1),
			domain.RoleAdmin,
			domain.TenantID(1),
			time.Duration(h.config.SessionDurationMinutes)*time.Minute,
		)
		if err != nil {
			http.Error(w, "Failed to create session", http.StatusInternalServerError)
			return
		}

		// Set secure cookie
		http.SetCookie(w, &http.Cookie{
			Name:     "session_id",
			Value:    sessionID,
			Path:     "/",
			MaxAge:   h.config.SessionDurationMinutes * 60,
			HttpOnly: true,
			Secure:   h.config.IsProduction(),
			SameSite: http.SameSiteLaxMode,
		})

		http.Redirect(w, r, "/admin", http.StatusFound)
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
	cookie, err := r.Cookie("session_id")
	if err == nil {
		h.store.Delete(cookie.Value)
	}

	// Clear cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   h.config.IsProduction(),
		SameSite: http.SameSiteLaxMode,
	})

	http.Redirect(w, r, "/login", http.StatusFound)
}
