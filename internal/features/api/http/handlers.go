package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Test(c *gin.Context) {
	c.String(http.StatusOK, `<p class="text-green-600 font-bold">HTMX is working!</p>`)
}

func (h *Handler) Confirm(c *gin.Context) {
	c.String(http.StatusOK, `<div class="alert alert-success"><span>Action confirmed!</span></div>`)
}

func (h *Handler) Greet(c *gin.Context) {
	name := c.PostForm("nameInput")
	if name == "" {
		name = "stranger"
	}
	c.String(http.StatusOK, "<p>Hello, <strong>"+name+"</strong>!</p>")
}
