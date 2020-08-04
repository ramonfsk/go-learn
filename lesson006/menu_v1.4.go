package main

import (
	"fmt"
	"os"
)

func main() {
	showIntro()
	showMenu()
	option := readOption()

	switch option {
	case 1:
		fmt.Println("\nMonitoring...")
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
