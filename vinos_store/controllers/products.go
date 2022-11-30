package controllers

import (
	"net/http"
	"server/models"
	"strconv"
	"text/template"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func GetIndex(w http.ResponseWriter, r *http.Request) {

	products := models.GetAllProducts()

	templates.ExecuteTemplate(w, "Index", products)
}

func GetNew(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "New", nil)
}

func InsertProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		priceAsFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			panic(err)
		}

		quantityAsInt, err := strconv.Atoi(quantity)
		if err != nil {
			panic(err)
		}

		models.InsertProduct(name, description, priceAsFloat, quantityAsInt)

	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")

	models.DeleteProduct(productId)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func EditProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")

	product := models.EditProduct(productId)

	templates.ExecuteTemplate(w, "Edit", product)

}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		priceAsFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			panic(err)
		}

		quantityAsInt, err := strconv.Atoi(quantity)
		if err != nil {
			panic(err)
		}

		idAsInt, err := strconv.Atoi(id)
		if err != nil {
			panic(err)
		}

		models.UpdateProduct(idAsInt, name, description, priceAsFloat, quantityAsInt)

	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)

}
