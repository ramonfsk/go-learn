package main

import (
	"fmt"
	"reflect"
)

func main() {
	name := "Amanda"
	age := 25
	version := 1.2 //default is float64
	fmt.Println("Hi, Srta.", name, "your age is", age)
	fmt.Println("This program is in version", version)

	fmt.Println("The type of var name is", reflect.TypeOf(name))
	fmt.Println("The type of var age is", reflect.TypeOf(age))
	fmt.Println("The type of var version is", reflect.TypeOf(version))
}
