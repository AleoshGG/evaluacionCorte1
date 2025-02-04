package controllers

import (
	"evaluacionCorte1/principal/db"
	"evaluacionCorte1/principal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	id_product, _ := strconv.ParseInt(id, 10, 64)

	newProducts := []models.Product{}

	for _, product := range db.Products {
    	if product.Id != int(id_product) {
        	newProducts = append(newProducts, product)
    	}
	}
	db.Products = newProducts

	c.JSON(http.StatusCreated, gin.H{
		"status": true,
	})
}