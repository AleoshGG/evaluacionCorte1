package controllers

import (
	"encoding/json"
	"evaluacionCorte1/replication/db"
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
	
	body, err := io.ReadAll(response.Body)

	if err != nil {
		fmt.Println("Error al leer el cuerpo de la respuesta:", err)
		return
	}

	// Deserializar el JSON en la estructura Response
	var res db.Response
	err = json.Unmarshal(body, &res)
	if err != nil {
		fmt.Println("Error al parsear JSON:", err)
		return
	}

	fmt.Println("Productos:", res.Producto)

	response.Body.Close()
	time.Sleep(5 * time.Second)
		
	}
	
}
