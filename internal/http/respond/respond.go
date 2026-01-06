package respond

import "github.com/gin-gonic/gin"

func HTML(c *gin.Context, code int, name string, data gin.H) {
	c.HTML(code, name, data)
}

func JSON(c *gin.Context, code int, data interface{}) {
	c.JSON(code, data)
}

func Error(c *gin.Context, code int, message string) {
	c.HTML(code, "error.tmpl", gin.H{
		"Title":   "Error",
		"Message": message,
		"Code":    code,
	})
}
