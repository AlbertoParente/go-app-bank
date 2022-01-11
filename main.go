package main

import "fmt"

type CurrentAccount struct{
	holder string
	agencyNumber int
	accountNumber int
	balance float64
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
	/*
	accountAlberto := CurrentAccount{holder: "Alberto Parente", agencyNumber: 123, accountNumber: 123456, balance: 150.00}
	accountJuliana := CurrentAccount{"Juliana", 321, 654321, 100.00}

	fmt.Println(accountAlberto, accountJuliana)

	var accountJulia *CurrentAccount
	accountJulia = new(CurrentAccount)
	accountJulia.holder = "Julia"
	accountJulia.agencyNumber = 147
	accountJulia.accountNumber = 1478963
	accountJulia.balance = 200.00

	fmt.Println(*accountJulia)
	*/

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

}
