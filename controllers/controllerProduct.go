package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/Golang-CRUD/entities"

	// "github.com/Golang-CRUD/src/entities"

	"github.com/Golang-CRUD/models"
)

func Index(response http.ResponseWriter, request *http.Request) {
	log.Println("Index")
	var productModel models.ModelProduct

	products, err := productModel.FindAll()
	if err != nil {
		log.Println(err)
		return
	}
	data := map[string]interface{}{
		"products": products,
	}

	temp, err := template.ParseFiles("views/product/index.html")
	if err != nil {
		log.Println(err)
		return
	}
	temp.Execute(response, data)

}

func ViewCreateProduct(w http.ResponseWriter, request *http.Request) {

	temp, err := template.ParseFiles("views/product/add.html")
	if err != nil {
		log.Println(err)
		return
	}
	temp.Execute(w, nil)
}

func CreateProduct(w http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	var product entities.Product
	product.Name = request.Form.Get("name")
	product.Price, _ = strconv.ParseFloat(request.Form.Get("price"), 64)
	product.Quantity, _ = strconv.ParseInt(request.Form.Get("quantity"), 10, 64)
	product.Description = request.Form.Get("description")
	var productModel models.ModelProduct
	productModel.AddData(&product)
	http.Redirect(w, request, "/products", http.StatusSeeOther)
}

func ViewUpdateProduct(w http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	var productModel models.ModelProduct

	products, errGetData := productModel.FindById(params["id"])
	if errGetData != nil {
		log.Println(errGetData)
		return
	}

	data := map[string]interface{}{
		"products": products,
	}

	temp, err := template.ParseFiles("views/product/update.html")
	if err != nil {
		log.Println(err)
		return
	}
	temp.Execute(w, data)
}

func UpdateProduct(w http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	id, _ := strconv.ParseInt(request.Form.Get("id"), 10, 64)

	var product entities.Product
	product.Name = request.Form.Get("name")
	product.Price, _ = strconv.ParseFloat(request.Form.Get("price"), 64)
	product.Quantity, _ = strconv.ParseInt(request.Form.Get("quantity"), 10, 64)
	product.Description = request.Form.Get("description")
	var productModel models.ModelProduct
	productModel.UpdateData(&product, id)
	http.Redirect(w, request, "/products", http.StatusSeeOther)
}

func DeleteProduct(w http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, _ := strconv.ParseInt(params["id"], 10, 64)

	var productModel models.ModelProduct
	productModel.DeleteData(id)
	http.Redirect(w, request, "/products", http.StatusSeeOther)
}
