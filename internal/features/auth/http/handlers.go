package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.tmpl", gin.H{
		"Title":        "Login",
		"ShowError":    false,
		"ErrorMessage": "",
		"Email":        "",
	})
}

func (h *Handler) LoginSubmit(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	if email == "admin@example.com" && password == "password" {
		c.Redirect(http.StatusFound, "/")
		return
	}

	c.HTML(http.StatusOK, "login.tmpl", gin.H{
		"Title":        "Login",
		"ShowError":    true,
		"ErrorMessage": "Invalid email or password",
		"Email":        email,
	})
}

func (h *Handler) ForgotPasswordPage(c *gin.Context) {
	c.HTML(http.StatusOK, "forgot-password.tmpl", gin.H{
		"Title": "Forgot Password",
	})
}

func (h *Handler) ForgotPasswordSubmit(c *gin.Context) {
	c.HTML(http.StatusOK, "forgot-password.tmpl", gin.H{
		"Title":   "Forgot Password",
		"Success": true,
		"Message": "If an account exists with this email, you will receive a password reset link shortly.",
	})
}

func (h *Handler) Logout(c *gin.Context) {
	c.Redirect(http.StatusFound, "/login")
}
