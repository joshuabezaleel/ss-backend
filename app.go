package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/joshuabezaleel/salestock-backend/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// App comment
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// Initialize comment
func (a *App) Initialize(user, password, dbname string) {
	connectionString := fmt.Sprintf("%s:%s@%s", user, password, dbname)

	var err error
	a.DB, err = sql.Open("mysql", connectionString)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Database connected")
	}
	a.Router = mux.NewRouter()
	a.InitializeRoutes()
}

// Run comment
func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(":8000", a.Router))
}

// InitializeRoutes comment
func (a *App) InitializeRoutes() {
	a.Router.HandleFunc("/products", a.GetProducts).Methods("GET")
	a.Router.HandleFunc("/product/{id}", a.GetProduct).Methods("GET")
	a.Router.HandleFunc("/product/{id}", a.CreateProduct).Methods("POST")
	a.Router.HandleFunc("/product/{id}", a.UpdateProduct).Methods("PUT")
	a.Router.HandleFunc("/product/{id}", a.DeleteProduct).Methods("DELETE")
}

// GetProducts comment
func (a *App) GetProducts(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))
	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}
	products, err := model.GetProducts(a.DB, start, count)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, products)
}

// GetProduct comment
func (a *App) GetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid Product ID")
		return
	}

	P := model.Product{ID: id}
	if err := P.GetProduct(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			RespondWithError(w, http.StatusNotFound, "Product not found")
		default:
			RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	RespondWithJSON(w, http.StatusOK, P)
}

// CreateProduct comment
func (a *App) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var P model.Product
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&P); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := P.CreateProduct(a.DB); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusCreated, P)

}

// UpdateProduct comment
func (a *App) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid Product ID")
		return
	}

	var P model.Product
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&P); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	P.ID = id

	if err := P.UpdateProduct(a.DB); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, P)
}

// DeleteProduct comment
func (a *App) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid Product ID")
		return
	}

	p := model.Product{ID: id}
	if err := p.DeleteProduct(a.DB); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

// RespondWithError comment
func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

// RespondWithJSON comment
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
