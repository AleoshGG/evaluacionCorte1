package routes

import (
	"evaluacionCorte1/principal/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	productsRouter := r.Group("/products") 
	{
		productsRouter.POST("/", controllers.CreateProduct)
		productsRouter.GET("/", controllers.GetProduct)
		productsRouter.PUT("/:id", controllers.UpdateProduct)
		productsRouter.DELETE("/:id", controllers.DeleteProduct)
	}
}