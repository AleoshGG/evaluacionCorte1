package controllers

import (
	"evaluacionCorte1/principal/db"
	"evaluacionCorte1/principal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	var newProduct models.Product

	id := c.Param("id")
	id_product, _ := strconv.ParseInt(id, 10, 64)
	
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error": "Datos inválidos: " + err.Error(),
		})
		return
	}

	for _, product := range db.Products {
    	if product.Id == int(id_product) {
        	product.Name = newProduct.Name
			product.CodeB = newProduct.CodeB
			product.Amount = newProduct.Amount
    	}
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"Producto": newProduct,
	})
}