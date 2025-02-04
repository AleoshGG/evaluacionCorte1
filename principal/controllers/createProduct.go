package controllers

import (
	"evaluacionCorte1/principal/db"
	"evaluacionCorte1/principal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	var product models.Product
	
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error": "Datos inválidos: " + err.Error(),
		})
		return
	}

	newProduct := models.NewProduct(product.Name, product.Amount, product.CodeB)

	db.Products = append(db.Products, *newProduct)

	c.JSON(http.StatusCreated, gin.H{
		"status": true,
		"Producto": newProduct,
	})

}