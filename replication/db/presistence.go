package db

import "evaluacionCorte1/replication/models"

type Response struct {
	Producto []models.Product `json:"Producto"`
	Status   bool             `json:"status"`
}
