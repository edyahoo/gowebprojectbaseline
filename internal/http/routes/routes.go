package routes

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"

	"goprojstructtest/internal/domain"
	adminhttp "goprojstructtest/internal/features/admin/http"
	authhttp "goprojstructtest/internal/features/auth/http"
	pageshttp "goprojstructtest/internal/features/pages/http"
	"goprojstructtest/internal/http/middleware"
	"goprojstructtest/internal/platform/config"
	"goprojstructtest/internal/platform/session"
	"goprojstructtest/internal/render"
)

func Setup(r chi.Router, logger *slog.Logger, renderer *render.Renderer, store *session.InMemoryStore, cfg *config.Config) {
	r.Use(middleware.Logger(logger))
	r.Use(middleware.Recovery(logger))

	// Serve static files
	fileServer := http.FileServer(http.Dir("./web/static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	// Public routes
	authHandler := authhttp.NewHandler(renderer, store, cfg)
	r.Get("/login", authHandler.LoginPage)
	r.Post("/login", authHandler.LoginSubmit)
	r.Get("/forgot-password", authHandler.ForgotPasswordPage)
	r.Post("/forgot-password", authHandler.ForgotPasswordSubmit)
	r.Post("/logout", authHandler.Logout)

	pageHandler := pageshttp.NewHandler(renderer)
	r.Get("/", authHandler.LoginPage)
	r.Get("/about", pageHandler.About)

	// Protected routes
	r.Group(func(r chi.Router) {
		r.Use(middleware.RequireAuth(store, logger))
		r.Get("/demo", pageHandler.Demo)
		r.Get("/index", pageHandler.Index)
	})

	// Admin routes - protected with admin role
	adminHandler := adminhttp.NewHandler(renderer)
	r.Group(func(r chi.Router) {
		r.Use(middleware.RequireAuth(store, logger))
		r.Use(middleware.RequireRole(domain.RoleAdmin, logger))
		r.Get("/admin", adminHandler.Dashboard)
		r.Get("/admin/", adminHandler.Dashboard)
	})
}
