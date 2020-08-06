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

func (c *CurrentAccount) Transfer(transferValue float64, transferAccount *CurrentAccount) bool {
	if transferValue > 0 && transferValue < c.balance {
		c.balance -= transferValue
		transferAccount.Deposit(transferValue)
		return true
	} else {
		return false
	}
}

func main() {
	account1 := CurrentAccount{name: "Amanda", balance: 300}
	account2 := CurrentAccount{name: "Ramon", balance: 0}

	status := account1.Transfer(100, &account2)
	fmt.Println(status)
	fmt.Println(account1)
	fmt.Println(account2)
}
