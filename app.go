package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user, password, dbname string) {
	connectionString := fmt.Sprintf("%s:%s@/%s", user, password, dbname)

	var err error
	a.DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	if err == nil {
		fmt.Println("Database connected")
	}
	a.Router = mux.NewRouter()
	a.InitializeRoutes()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(":8000", a.Router))
}

func (a *App) InitializeRoutes() {
	a.Router.HandleFunc("/products", GetProducts).Methods("GET")
	a.Router.HandleFunc("/product/{id}", GetProduct).Methods("GET")
	a.Router.HandleFunc("/product/{id}", CreateProduct).Methods("POST")
	a.Router.HandleFunc("/product/{id}", DeleteProduct).Methods("DELETE")
}
