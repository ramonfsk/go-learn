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

	switch option {
	case 1:
		fmt.Println("\nMonitoring...")
	case 2:
		fmt.Println("\nShowing logs...")
	case 0:
		fmt.Println("\nExit out program...")
	default:
		fmt.Println("\nInstruction invalid!")
	}
}
