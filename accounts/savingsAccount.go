package accounts

import "github.com/go-app-bank/customers"

type SavingsAccount struct {
	Holder customers.Holder
	AgencyNumber int
	AccountNumber int
	Operation int
	balance float64
}

func (c *SavingsAccount) Deposit(depositAmount float64) (string, float64) {
	if depositAmount > 0 {
		c.balance += depositAmount
		return "Deposit made successfully!", c.balance
	} else {
		return "Deposit amount less than zero!", c.balance
	}
}

func (c *SavingsAccount) ToWithdraw(withdrawValue float64) string {
	flagWithdraw := withdrawValue > 0 && withdrawValue <= c.balance

	if flagWithdraw {
		c.balance -= withdrawValue
		return "Withdrawal successful!"
	} else {
		return "Insufficient funds!"
	}
}

func (c *SavingsAccount) GetBalance() float64 {
    return c.balance
}

