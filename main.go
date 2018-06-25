package main

var products []Product

func main() {
	var dbUsername, dbPassword, dbName string
	dbUsername = "root"
	dbPassword = ""
	dbName = "tcp(localhost:8080)/salestock_backend"
	products = append(products, Product{ID: "1", Name: "Phone", Price: "12345"})
	products = append(products, Product{ID: "2", Name: "Laptop", Price: "12345678"})
	a := App{}
	a.Initialize(dbUsername, dbPassword, dbName)
	a.Run(":8000")
}
