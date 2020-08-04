package main

import "fmt"

func main() {
	name := "Amanda"
	version := 1.0

	fmt.Println("Hi Ms.", name)
	fmt.Println("This program is in version", version)

	fmt.Println("\n1 - Start monitor")
	fmt.Println("2 - Show logs")
	fmt.Println("0 - Exit")

	var option int
	fmt.Scan(&option)
	fmt.Println("\nThis var option is allocated in address", &option)
}
