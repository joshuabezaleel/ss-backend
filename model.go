package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Product struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Price string `json:"price,omitempty"`
}

// GetProducts function description
func GetProducts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(products)
	fmt.Println("/products endpoint was hit")
	fmt.Println("Retrieving all the products")
}

// GetProduct function description
func GetProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, product := range products {
		if product.ID == params["id"] {
			json.NewEncoder(w).Encode(product)
			return
		}
	}
	json.NewEncoder(w).Encode(&Product{})
	fmt.Println("GET /product/{id} endpoint was hit")
	fmt.Println("Retrieving one particular product")
}

// CreateProduct function description
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var product Product
	_ = json.NewDecoder(r.Body).Decode(&product)
	product.ID = params["id"]
	products = append(products, product)
	json.NewEncoder(w).Encode(products)
	fmt.Println("POST /product/{id} endpoint was hit")
	fmt.Println("Creating a new product")
}

// DeleteProduct function description
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, product := range products {
		if product.ID == params["id"] {
			products = append(products[:index], products[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(products)
	fmt.Println("DELETE /product/{id} endpoint was hit")
	fmt.Println("Deleting a product")
}
