package model

import (
	"database/sql"
	"fmt"
)

type Product struct {
	ID    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Price string `json:"price,omitempty"`
}

// GetProducts function description
func GetProducts(db *sql.DB, start, count int) ([]Product, error) {
	statement := fmt.Sprintf("SELECT id, name, price FROM products LIMIT %d OFFSET %d", count, start)
	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	Products := []Product{}

	for rows.Next() {
		var P Product
		if err := rows.Scan(&P.ID, &P.Name, &P.Price); err != nil {
			return nil, err
		}
		Products = append(Products, P)
	}

	return Products, nil
	// json.NewEncoder(w).Encode(products)
	// fmt.Println("/products endpoint was hit")
	// fmt.Println("Retrieving all the products")
}

// GetProduct function description
func (P *Product) GetProduct(db *sql.DB) error {
	return db.QueryRow("SELECT id, name, price FROM products WHERE id=?", P.ID).Scan(&P.ID, &P.Name, &P.Price)
	// params := mux.Vars(r)
	// for _, product := range products {
	// 	if product.ID == params["id"] {
	// 		json.NewEncoder(w).Encode(product)
	// 		return
	// 	}
	// }
	// json.NewEncoder(w).Encode(&Product{})
	// fmt.Println("GET /product/{id} endpoint was hit")
	// fmt.Println("Retrieving one particular product")
}

// CreateProduct function description
func (P *Product) CreateProduct(db *sql.DB) error {
	statement := fmt.Sprintf("INSERT INTO products (name, price) VALUES ('%s','%s')", P.Name, P.Price)
	_, err := db.Exec(statement)
	// err := db.QueryRow("INSERT INTO products(name, price) VALUES ('?','?') RETURNING id", P.Name, P.Price).Scan(&P.ID)

	if err != nil {
		return err
	}

	err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&P.ID)

	if err != nil {
		return err
	}

	return nil
	// params := mux.Vars(r)
	// var product Product
	// _ = json.NewDecoder(r.Body).Decode(&product)
	// product.ID = params["id"]
	// products = append(products, product)
	// json.NewEncoder(w).Encode(products)
	// fmt.Println("POST /product/{id} endpoint was hit")
	// fmt.Println("Creating a new product")
}

// DeleteProduct function description
func (P *Product) DeleteProduct(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM products WHERE id=%d", P.ID)

	return err
	// params := mux.Vars(r)
	// for index, product := range products {
	// 	if product.ID == params["id"] {
	// 		products = append(products[:index], products[index+1:]...)
	// 		break
	// 	}
	// }
	// json.NewEncoder(w).Encode(products)
	// fmt.Println("DELETE /product/{id} endpoint was hit")
	// fmt.Println("Deleting a product")
}
