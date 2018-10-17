package main

const (
	host     = "localhost"
	port     = 5432
	user     = "joshuabezaleel"
	password = "neph224301"
  )

func main() {
	// products = append(products, Product{ID: 1, Name: "Phone", Price: "12345"})
	// products = append(products, Product{ID: 2, Name: "Laptop", Price: "12345678"})
	dbname := "salestock_backend"
	a := App{}
	a.Initialize(host, port, user, password, dbname)
	a.Run(":8000")
}
