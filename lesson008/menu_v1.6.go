package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitorCount = 3
const delay = 5

func main() {
	showIntro()

	for {
		showMenu()
		option := readOption()

		switch option {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("\nShowing logs...")
		case 0:
			fmt.Println("\nExit out program...")
			os.Exit(0)
		default:
			fmt.Println("\nInstruction invalid!")
			os.Exit(-1)
		}
	}
}

func showIntro() {
	name := "Amanda"
	version := 1.0

	fmt.Println("Hi Ms.", name)
	fmt.Println("This program is in version", version)
}

func showMenu() {
	fmt.Println("\n1 - Start monitor")
	fmt.Println("2 - Show logs")
	fmt.Println("0 - Exit")
}

func readOption() int {
	var option int
	fmt.Scan(&option)
	return option
}

func startMonitoring() {
	fmt.Println("\nMonitoring...")
	urls := []string{"https://www.ibm.com", "https://www.poder360.com.br", "https://random-status-code.herokuapp.com"}

	for i := 0; i < monitorCount; i++ {
		for _, url := range urls {
			getStatusURL(url)
		}
		time.Sleep(delay * time.Second)
		fmt.Println()
	}
}

func getStatusURL(url string) {
	response, _ := http.Get(url)
	if response.StatusCode == 200 {
		fmt.Println("Url:", url, "is OK!")
	} else {
		fmt.Println("Url:", url, "isn't OK, status code:", response.StatusCode)
	}
}
