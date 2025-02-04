package controllers

import (
	"evaluacionCorte1/principal/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProduct(c *gin.Context) {

	c.JSON(http.StatusCreated, gin.H{
		"status": true,
		"Producto": db.Products,
	})
}