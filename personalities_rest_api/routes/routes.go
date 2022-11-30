package routes

import (
	"log"
	"net/http"
	"personalities_api/controllers"
	"personalities_api/middleware"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {

	r := mux.NewRouter()

	r.Use(middleware.ContentTypeMiddleware)

	r.HandleFunc("/", controllers.GetHome)

	r.HandleFunc(
		"/api/personalities",
		controllers.GetAllPersonalities,
	).Methods(http.MethodGet)

	r.HandleFunc(
		"/api/personalities/{id}",
		controllers.GetPersonalityById,
	).Methods(http.MethodGet)

	r.HandleFunc(
		"/api/personalities",
		controllers.CreatePersonality,
	).Methods(http.MethodPost)

	r.HandleFunc(
		"/api/personalities/{id}",
		controllers.DeletePersonality,
	).Methods(http.MethodDelete)

	r.HandleFunc(
		"/api/personalities/{id}",
		controllers.EditPersonality,
	).Methods(http.MethodPut)

	log.Fatal(
		http.ListenAndServe(
			":8000",
			handlers.CORS(
				handlers.AllowedOrigins(
					[]string{"*"}))(r),
		),
	)

}
