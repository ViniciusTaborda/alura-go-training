package main

import (
	"fmt"
)

func main() {
	var user_name string
	// Variable with no type? Type inference!!
	program_version := "0.0.1"

	fmt.Scanf("%s", &user_name)

	fmt.Println("Hello, Sr.", user_name)
	fmt.Println("Version:", program_version)
}
