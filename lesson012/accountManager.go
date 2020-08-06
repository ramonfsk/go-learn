package main

import "fmt"

type CurrentAccount struct {
	name          string
	agencyNumber  int
	accountNumber int
	balance       float64
}

func (c *CurrentAccount) withdraw(valueWithdraw float64) string {
	canWithdraw := valueWithdraw > 0 && valueWithdraw <= c.balance
	if canWithdraw == true {
		c.balance -= valueWithdraw
		return "Successful withdrawal!"
	} else {
		return "Insufficient balance..."
	}
}

func main() {
	account := CurrentAccount{}
	account.name = "Amanda"
	account.balance = 500

	fmt.Println(account)
	fmt.Println(account.withdraw(-200))
	fmt.Println(account)
}
