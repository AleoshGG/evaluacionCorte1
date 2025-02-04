package models

type Product struct {
	Id     int
	Name   string
	Amount int
	CodeB  string
}

var count = 0

func NewProduct(name string, amount int, codeB string) *Product {
	count++
	return &Product{Id: count, Name: name, Amount: amount, CodeB: codeB}
}