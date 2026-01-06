package routes

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	
	authhttp "goprojstructtest/internal/features/auth/http"
	pageshttp "goprojstructtest/internal/features/pages/http"
	"goprojstructtest/internal/http/middleware"
)

func Setup(r *gin.Engine, logger *slog.Logger) {
	r.Use(middleware.Logger(logger))
	r.Use(middleware.Recovery(logger))

	r.Static("/static", "./web/static")

	public := r.Group("")
	{
		authHandler := authhttp.NewHandler()
		public.GET("/login", authHandler.LoginPage)
		public.POST("/login", authHandler.LoginSubmit)
		public.GET("/forgot-password", authHandler.ForgotPasswordPage)
		public.POST("/forgot-password", authHandler.ForgotPasswordSubmit)
		public.POST("/logout", authHandler.Logout)

		pageHandler := pageshttp.NewHandler()
		public.GET("/", pageHandler.Index)
		public.GET("/about", pageHandler.About)
		public.GET("/demo", pageHandler.Demo)
	}

	protected := r.Group("")
	protected.Use(middleware.RequireAuth())
	{
	}
}
