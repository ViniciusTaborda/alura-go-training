package main

import (
	"fmt"
	"net/http"
	"server/routes"
)

func main() {

	port := "8000"

	routes.LoadRoutes()

	fmt.Printf("Listening on port %s...", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)

	if err != nil {
		panic(err)
	}

}
