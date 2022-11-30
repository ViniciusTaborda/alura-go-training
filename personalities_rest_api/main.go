package main

import (
	"fmt"
	"personalities_api/routes"
)

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	fmt.Println("Starting up rest server...")
	routes.HandleRequest()

}
