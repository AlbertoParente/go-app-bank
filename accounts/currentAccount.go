package accounts

import "github.com/go-app-bank/customers"

type CurrentAccount struct {
	Holder        customers.Holder
	AgencyNumber  int
	AccountNumber int
	balance       float64
}

func (c *CurrentAccount) Deposit(depositAmount float64) (string, float64) {
	if depositAmount > 0 {
		c.balance += depositAmount
		return "Deposit made successfully!", c.balance
	} else {
		return "Deposit amount less than zero!", c.balance
	}
}

func (c *CurrentAccount) ToWithdraw(withdrawValue float64) string {
	flagWithdraw := withdrawValue > 0 && withdrawValue <= c.balance

	if flagWithdraw {
		c.balance -= withdrawValue
		return "Withdrawal successful!"
	} else {
		return "Insufficient funds!"
	}
}

func (c *CurrentAccount) Transfer(transferValue float64, accountDestination *CurrentAccount) bool {
	if transferValue < c.balance && transferValue > 0 {
		c.balance -= transferValue
		accountDestination.ToWithdraw(transferValue)
		return true
	} else {
		return false
	}
}

func (c *CurrentAccount) GetBalance() float64 {
    return c.balance
}
