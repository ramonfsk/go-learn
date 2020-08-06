package main

import "fmt"

type CurrentAccount struct {
	name          string
	agencyNumber  int
	accountNumber int
	balance       float64
}

func (c *CurrentAccount) Withdraw(withdrawValue float64) string {
	canWithdraw := withdrawValue > 0 && withdrawValue <= c.balance
	if canWithdraw == true {
		c.balance -= withdrawValue
		return "Successful withdrawal!"
	} else {
		return "Insufficient balance..."
	}
}

func (c *CurrentAccount) Deposit(depositValue float64) (string, float64) {
	if depositValue > 0 {
		c.balance += depositValue
		return "Deposit successfully!", c.balance
	} else {
		return "Deposit value is invalid!", c.balance
	}
}

func main() {
	account := CurrentAccount{}
	account.name = "Amanda"
	account.balance = 500

	// fmt.Println(account.Withdraw(-200))
	// fmt.Println(account)
	status, value := account.Deposit(1000)
	fmt.Println(status, value)
}
