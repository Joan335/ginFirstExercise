package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"time"
)

func SetupRouter(r *gin.Engine) {
	r.LoadHTMLGlob("templates/*")
	
	r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/api/time", func(c *gin.Context) {
		now := time.Now().Format("Monday, 02 January 2006 15:04:05")
		c.JSON(http.StatusOK, gin.H{"time": now})
	})
}

