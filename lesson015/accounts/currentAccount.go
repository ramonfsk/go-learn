package accounts

type CurrentAccount struct {
	Name          string
	AgencyNumber  int
	AccountNumber int
	Balance       float64
}

func (c *CurrentAccount) Withdraw(withdrawValue float64) string {
	canWithdraw := withdrawValue > 0 && withdrawValue <= c.Balance
	if canWithdraw == true {
		c.Balance -= withdrawValue
		return "Successful withdrawal!"
	} else {
		return "Insufficient balance..."
	}
}

func (c *CurrentAccount) Deposit(depositValue float64) (string, float64) {
	if depositValue > 0 {
		c.Balance += depositValue
		return "Deposit successfully!", c.Balance
	} else {
		return "Deposit value is invalid!", c.Balance
	}
}

func (c *CurrentAccount) Transfer(transferValue float64, transferAccount *CurrentAccount) bool {
	if transferValue > 0 && transferValue < c.Balance {
		c.Balance -= transferValue
		transferAccount.Deposit(transferValue)
		return true
	} else {
		return false
	}
}
