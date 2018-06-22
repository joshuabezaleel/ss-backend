package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Product type description
type Product struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Price string `json:"price,omitempty"`
}

var products []Product

func main() {
	products = append(products, Product{ID: "1", Name: "Phone", Price: "12345"})
	products = append(products, Product{ID: "2", Name: "Laptop", Price: "12345678"})
	router := mux.NewRouter()
	router.HandleFunc("/products", GetProducts).Methods("GET")
	router.HandleFunc("/product/{id}", GetProduct).Methods("GET")
	router.HandleFunc("/product/{id}", CreateProduct).Methods("POST")
	router.HandleFunc("/product/{id}", DeleteProduct).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
	fmt.Println("Salestock-backend-API is up and running...")
	// fmt.Println("Hello world")
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
		json.NewEncoder(w).Encode(products)
	}
	fmt.Println("DELETE /product/{id} endpoint was hit")
	fmt.Println("Deleting a product")
}
