package main

import (
	"evaluacionCorte1/replication/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.RegisterRoutes(r)

	port := ":8081"
	r.Run(port)
}