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
			showLogs()
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
	urls := readUrlsFile()

	for i := 0; i < monitorCount; i++ {
		for _, url := range urls {
			getStatusURL(url)
		}
		time.Sleep(delay * time.Second)
		fmt.Println()
	}
}

func getStatusURL(url string) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("An error occurred, details here", err)
	}

	if res.StatusCode == 200 {
		fmt.Println("Url:", url, "is OK!")
		writeLog(url, true)
	} else {
		fmt.Println("Url:", url, "isn't OK, status code:", res.StatusCode)
		writeLog(url, false)
	}
}

func readUrlsFile() []string {
	var urls []string

	file, err := os.Open("urls.txt")
	// file, err := ioutil.ReadFile("urls.txt")
	if err != nil {
		fmt.Println("An error occurred, details here:", err)
	}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		urls = append(urls, line)
		if err == io.EOF {
			break
		}
	}

	return urls
}

func writeLog(site string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("An error occurred, details here:", err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	file.Close()
}

func showLogs() {
	fmt.Println("\nShowing logs...")

	file, err := ioutil.ReadFile("log.txt")
	if err != nil {
		fmt.Println("An error occurred, details here:", err)
	}

	fmt.Println(string(file))
}
