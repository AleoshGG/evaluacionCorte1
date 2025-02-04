package main

import (
	"evaluacionCorte1/principal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()


	routes.RegisterRoutes(r)

	r.Run()
}