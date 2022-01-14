package accounts

type CurrentAccount struct {
	Holder        string
	AgencyNumber  int
	AccountNumber int
	Balance       float64
}

func (c *CurrentAccount) ToWithdraw(withdrawValue float64) string {
	flagWithdraw := withdrawValue > 0 && withdrawValue <= c.Balance

	if flagWithdraw {
		c.Balance -= withdrawValue
		return "Withdrawal successful!"
	} else {
		return "Insufficient funds!"
	}
}

func (c *CurrentAccount) Deposit(depositAmount float64) (string, float64) {
	if depositAmount > 0 {
		c.Balance += depositAmount
		return "Deposit made successfully!", c.Balance
	} else {
		return "Deposit amount less than zero!", c.Balance
	}
}

func (c *CurrentAccount) Transfer(transferValue float64, accountDestination *CurrentAccount) bool {
	if transferValue < c.Balance && transferValue > 0 {
		c.Balance -= transferValue
		accountDestination.ToWithdraw(transferValue)
		return true
	} else {
		return false
	}
}
