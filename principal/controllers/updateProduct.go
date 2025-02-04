package controllers

import (
	"evaluacionCorte1/principal/db"
	"evaluacionCorte1/principal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateProduct(c *gin.Context) {
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

	// Recorrer el slice por índice para modificarlo directamente
	for i := range db.Products {
		if db.Products[i].Id == int(id_product) {
			// Modificar directamente el producto en la lista
			db.Products[i].Name = newProduct.Name
			db.Products[i].CodeB = newProduct.CodeB
			db.Products[i].Amount = newProduct.Amount

			// Responder con el producto actualizado
			c.JSON(http.StatusOK, gin.H{
				"status":   true,
				"Producto": db.Products[i],
			})
			return
		}
	}

	// Si no se encuentra el producto, responder con error
	c.JSON(http.StatusNotFound, gin.H{
		"status": false,
		"error":  "Producto no encontrado",
	})
}