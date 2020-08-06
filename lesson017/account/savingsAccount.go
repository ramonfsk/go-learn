package account

import cl "../client"

type SavingsAccount struct {
	Holder                                 cl.Client
	AgencyNumber, AccountNumber, Operation int
	balance                                float64
}

func (c *SavingsAccount) Withdraw(withdrawValue float64) string {
	canWithdraw := withdrawValue > 0 && withdrawValue <= c.balance
	if canWithdraw == true {
		c.balance -= withdrawValue
		return "Successful withdrawal!"
	} else {
		return "Insufficient balance..."
	}
}

func (c *SavingsAccount) Deposit(depositValue float64) (string, float64) {
	if depositValue > 0 {
		c.balance += depositValue
		return "Deposit successfully!", c.balance
	} else {
		return "Deposit value is invalid!", c.balance
	}
}

func (c *SavingsAccount) ShowBalance() float64 {
	return c.balance
}
