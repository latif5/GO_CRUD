package main

import (
	//package dari go
	"log"
	"net/http"

	//package dari project kita
	"github.com/Golang-CRUD/controllers"
	//package dari luar
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/products", controllers.Index).Methods("GET")
	r.HandleFunc("/products/add", controllers.ViewCreateProduct).Methods("GET")
	r.HandleFunc("/products/add", controllers.CreateProduct).Methods("POST")
	// r.HandleFunc("/products/{id}", controllers.getProduct).Methods("GET")
	r.HandleFunc("/products/update/{id}", controllers.ViewUpdateProduct).Methods("GET")
	r.HandleFunc("/products/update", controllers.UpdateProduct).Methods("PUT")
	r.HandleFunc("/products/delete/{id}", controllers.DeleteProduct).Methods("GET")
	// http.HandleFunc("/products", controllers.Index)

	log.Fatal(http.ListenAndServe(":8080", r))
}
