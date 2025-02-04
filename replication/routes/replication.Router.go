package routes

import (
	cR "evaluacionCorte1/replication/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	productsRouter := r.Group("/products") 
	{
		productsRouter.GET("/", cR.Check)
	}
}