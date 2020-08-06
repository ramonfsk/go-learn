package main

import (
	"fmt"

	c "./account"
)

func main() {
	// client1 := cl.Client{"Amanda", "000.000.000-99", "Redatora"}
	account1 := c.CurrentAccount{}
	account1.Deposit(100)

	fmt.Println(account1)
}
