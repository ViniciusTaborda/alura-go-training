package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const numberOfMonitors = 5
const waitForMonitorInSeconds = 5
const urlsFileName = "urls.txt"
const logsFileName = "logs.txt"

func main() {

	//Infinite loop
	for {
		printIntro()

		switch getUserOption() {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Searching for logs..")
			showLogs()
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

	fmt.Println("")

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

	_, err := fmt.Scan(&optionVar)

	if err != nil {
		return -1
	}

	return optionVar
}

func startMonitoring() {
	fmt.Println("Starting monitoring...")
	fmt.Println("")

	urlsToMonitor := getUrlsFromFile()

	for i := 0; i < numberOfMonitors; i++ {
		for _, url := range urlsToMonitor {
			testWebsite(url)
		}
		time.Sleep(waitForMonitorInSeconds * time.Second)
		fmt.Println("")
	}

}

func testWebsite(url string) {
	response, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	if response.StatusCode == 200 {
		fmt.Printf("%s is up and running!", url)
		saveLogs(url, false)
	} else {
		fmt.Printf("%s is down and returned %s!", url, response.Status)
		saveLogs(url, false)
	}

	fmt.Println("")
}

func getUrlsFromFile() []string {

	var urls []string

	file, err := os.Open(urlsFileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	if err != nil {
		fmt.Println(err)
	}

	for {

		line, err := reader.ReadString('\n')
		urls = append(urls, strings.TrimSpace(line))

		if err == io.EOF {
			break
		}

	}

	return urls
}

func saveLogs(url string, isOnline bool) {
	file, err := os.OpenFile(logsFileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	defer file.Close()

	currentTime := time.Now().Format(time.RFC3339Nano)

	file.WriteString(
		fmt.Sprintf(
			"%s - %s - Online:  %s\n", currentTime, url, strconv.FormatBool(isOnline),
		),
	)
}

func showLogs() {
	file, err := ioutil.ReadFile(logsFileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	fmt.Println(string(file))

}
