package routes

import (
	"net/http"
	"server/controllers"
)

func LoadRoutes() {

	http.HandleFunc("/", controllers.GetIndex)
	http.HandleFunc("/new", controllers.GetNew)
	http.HandleFunc("/insert", controllers.InsertProduct)
	http.HandleFunc("/delete", controllers.DeleteProduct)
	http.HandleFunc("/edit", controllers.EditProduct)
	http.HandleFunc("/update", controllers.UpdateProduct)

}
