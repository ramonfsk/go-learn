package account

import c "../client"

type CurrentAccount struct {
	Holder        c.Client
	AgencyNumber  int
	AccountNumber int
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

func (c *CurrentAccount) showBalance() float64 {
	return c.balance
}
