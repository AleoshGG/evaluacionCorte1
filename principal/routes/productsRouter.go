package routes

import (
	cP "evaluacionCorte1/principal/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	productsRouter := r.Group("/products") 
	{
		productsRouter.POST("/", cP.CreateProduct)
		productsRouter.GET("/", cP.GetProduct)
		productsRouter.PUT("/:id", cP.UpdateProduct)
		productsRouter.DELETE("/:id", cP.DeleteProduct)
	}
}