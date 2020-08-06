package main

import (
	"fmt"

	c "./accounts"
)

func main() {
	account1 := c.CurrentAccount{Name: "Amanda", Balance: 300}
	account2 := c.CurrentAccount{Name: "Ramon", Balance: 0}

	status := account1.Transfer(100, &account2)
	fmt.Println(status)
	fmt.Println(account1)
	fmt.Println(account2)
}
