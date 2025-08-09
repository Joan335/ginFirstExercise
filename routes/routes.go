package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	r.LoadHTMLGlob("templates/*")
	
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"message": "Hola Mundo"})
	})
}