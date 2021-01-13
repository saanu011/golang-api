package handler

import (
	"fmt"
	"strconv"
)

type Product struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var productData = []Product{
	{ID: 1, Name: "Mouse"},
	{ID: 2, Name: "Keyboard"},
	{ID: 3, Name: "Monitor"},
	{ID: 4, Name: "Joystick"},
	{ID: 5, Name: "Webcam"},
}

func (p Product) createProduct(product Product) error {

	// next available id for product
	index := nextId(productData)
	product.ID = index

	// add product to the product list
	productData = append(productData, product)
	return nil
}

func (p Product) getProduct(id string) (Product, error) {

	// get product's index from the product list through id
	index := indexByID(productData, id)
	if index < 0 {
		return Product{}, fmt.Errorf("no product found")
	}
	return productData[index], nil
}

func (p Product) updateProduct(data Product, id string) error {

	// get product's index from the product list through id
	index := indexByID(productData, id)
	if index < 0 {
		return fmt.Errorf("no product found")
	}
	productData[index] = data
	return nil
}

func (p Product) deleteProduct(id string) error {

	// get product's index from the product list through id
	index := indexByID(productData, id)
	if index < 0 {
		return fmt.Errorf("no product found")
	}
	// remove the product
	productData = append(productData[:index], productData[index+1:]...)
	return nil
}

func nextId(products []Product) int {
	var nextId int
	for i := 0; i < len(products); i++ {
		if products[i].ID > nextId {
			nextId = products[i].ID
		}
	}
	return nextId
}

func indexByID(products []Product, idString string) int {

	id := stringToInt(idString)
	for i := 0; i < len(products); i++ {
		if products[i].ID == id {
			return i
		}
	}
	return -1
}

func stringToInt(idString string) int {

	id, err := strconv.Atoi(idString)
	if err != nil {
		return 0
	}
	return id
}
