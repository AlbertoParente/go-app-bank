package main

import "fmt"

type CurrentAccount struct {
	holder        string
	agencyNumber  int
	accountNumber int
	balance       float64
}

func (c *CurrentAccount) toWithdraw(withdrawValue float64) string {
	flagWithdraw := withdrawValue > 0 && withdrawValue <= c.balance

	if flagWithdraw {
		c.balance -= withdrawValue
		return "Withdrawal successful!"
	} else {
		return "Insufficient funds!"
	}
}

func main() {
	accountAlberto := CurrentAccount{}
	accountAlberto.holder = "Alberto Parente"
	accountAlberto.agencyNumber = 123
	accountAlberto.accountNumber = 123456789
	accountAlberto.balance = 500

	fmt.Println(accountAlberto.balance)
	fmt.Println(accountAlberto.toWithdraw(300))
	fmt.Println(accountAlberto.balance)
	fmt.Println(accountAlberto.toWithdraw(500))
	fmt.Println(accountAlberto.balance)
	fmt.Println(accountAlberto.toWithdraw(-300))
	fmt.Println(accountAlberto.balance)

	fmt.Println(accountAlberto.balance)
	status, value := accountAlberto.Deposit(2000)
	fmt.Println(status, value)

    accountJuliana := CurrentAccount{holder:"Gustavo", balance: 100}

    fmt.Pirintln(accountJuliana)
    fmt.Pirintln(accountAlberto)

	status := accountAlberto.Transfer(200, accountJuliana)
}

func (c *CurrentAccount) Deposit(depositAmount float64) (string, float64) {
	if depositAmount > 0 {
		c.balance += depositAmount
		return "Deposit made successfully!", c.balance
	} else {
		return "Deposit amount less than zero!", c.balance
	}
}

func (c *CurrentAccount) Transfer(transferValue float64, accountDestination CurrentAccount) bool {
    if transferValue < c.balance {
        c.balance -= transferValue 
        accountDestination.toWithdraw(transferValue)
        return true
    } else {
        return false
    }
}
