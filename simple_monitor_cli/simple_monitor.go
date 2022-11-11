package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	for {
		printIntro()

		switch getUserOption() {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Searching for logs..")
		case 0:
			fmt.Println("Exiting...")
			// This function informs that the program exited correctly;
			os.Exit(0)
		default:
			fmt.Println("This is not a valid option...")
			// This function informs that the program did not exited correctly;
			os.Exit(-1)
		}
	}
}

func printIntro() {

	fmt.Println("Welcome to the SimMon CLI (A simple monitor CLI)!")

	fmt.Println("")

	fmt.Println("To start choose an option: ")
	fmt.Println("1 - Start monitoring")
	fmt.Println("2 - Show logs")
	fmt.Println("0 - Exit")

	fmt.Println("")

}

func getUserOption() int {

	var optionVar int

	returnedOption, err := fmt.Scan(&optionVar)

	if err != nil {
		return -1
	}

	return returnedOption
}

func startMonitoring() {
	fmt.Println("Starting monitoring...")

	urlToMonitor := "http://random-status-code.herokuapp.com/"

	response, err := http.Get(urlToMonitor)

	if err != nil {
		fmt.Println(err)
	}

	if response.StatusCode == 200 {
		fmt.Printf("%s is up and running!", urlToMonitor)
	} else {
		fmt.Printf("%s is down and returned %s!", urlToMonitor, response.Status)
	}

}
