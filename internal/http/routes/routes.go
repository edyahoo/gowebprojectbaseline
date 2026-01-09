package routes

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"

	authhttp "goprojstructtest/internal/features/auth/http"
	pageshttp "goprojstructtest/internal/features/pages/http"
	"goprojstructtest/internal/http/middleware"
	"goprojstructtest/internal/render"
)

func Setup(r chi.Router, logger *slog.Logger, renderer *render.Renderer) {
	r.Use(middleware.Logger(logger))
	r.Use(middleware.Recovery(logger))

	// Serve static files
	fileServer := http.FileServer(http.Dir("./web/static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	// Public routes
	authHandler := authhttp.NewHandler(renderer)
	r.Get("/login", authHandler.LoginPage)
	r.Post("/login", authHandler.LoginSubmit)
	r.Get("/forgot-password", authHandler.ForgotPasswordPage)
	r.Post("/forgot-password", authHandler.ForgotPasswordSubmit)
	r.Post("/logout", authHandler.Logout)

	pageHandler := pageshttp.NewHandler(renderer)
	r.Get("/", authHandler.LoginPage)
	r.Get("/about", pageHandler.About)
	r.Get("/demo", pageHandler.Demo)
	r.Get("/index", pageHandler.Index)

	// Protected routes
	r.Group(func(r chi.Router) {
		r.Use(middleware.RequireAuth())
		// Add protected routes here
	})
}
