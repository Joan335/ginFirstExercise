package main

import (
	"ejercicio1/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.SetupRouter(r)

	r.Run(":8080")
}