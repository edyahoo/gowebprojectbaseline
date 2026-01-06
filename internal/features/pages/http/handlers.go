package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"Title": "Home",
	})
}

func (h *Handler) About(c *gin.Context) {
	c.HTML(http.StatusOK, "about.tmpl", gin.H{
		"Title": "About",
	})
}

func (h *Handler) Demo(c *gin.Context) {
	c.HTML(http.StatusOK, "demo.tmpl", gin.H{
		"Title": "Demo",
	})
}
