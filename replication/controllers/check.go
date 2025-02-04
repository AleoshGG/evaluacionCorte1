package controllers

import (
	"encoding/json"
	"evaluacionCorte1/replication/models"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Check(c *gin.Context) {
	for  {
		response, err := http.Get("http://localhost:8080/products/")

	if err != nil {
		fmt.Println("Error al hacer la petición:", err)
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	// Deserializar el JSON en un array de Product
	var products []models.Product
	err = json.Unmarshal(body, &products)
	if err != nil {
		fmt.Println("Error al parsear JSON:", err)
		return
	}

	// Mostrar los productos obtenidos
	fmt.Println("Productos obtenidos:")
	for _, product := range products {
		fmt.Printf("ID: %d, Nombre: %s, Cantidad: %d, Código: %s\n",
			product.Id, product.Name, product.Amount, product.CodeB)
	}

	time.Sleep(5 * time.Second)
		
	}
	
}
