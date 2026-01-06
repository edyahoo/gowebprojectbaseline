package http

import "github.com/gin-gonic/gin"

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) List(c *gin.Context) {
}

func (h *Handler) Create(c *gin.Context) {
}

func (h *Handler) Update(c *gin.Context) {
}

func (h *Handler) Delete(c *gin.Context) {
}
