package main

import (
	"fmt"

	c "./account"
)

type verifyAccount interface {
	Withdraw(value float64) string
}

func PayTicket(account verifyAccount, ticketValue float64) {
	account.Withdraw(ticketValue)
}

func main() {
	account1 := c.SavingsAccount{}
	account1.Deposit(100)
	PayTicket(&account1, 60)

	fmt.Println(account1.ShowBalance())
}
