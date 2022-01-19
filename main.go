package main

import (
	"fmt"
	"github.com/go-app-bank/accounts"
)

func PayBankingBillet(account checkAccount, valueBankingBillet float64) {
	account.ToWithdraw(valueBankingBillet)
}

type checkAccount interface {
	ToWithdraw(value float64) string
}

func main() {

	accountAlberto := accounts.SavingsAccount{}
	accountAlberto.Deposit(1000)
	accountAlberto.ToWithdraw(100)
	PayBankingBillet(&accountAlberto, 200)
	fmt.Println(accountAlberto)
	
	accountJuliana := accounts.SavingsAccount{}
	accountJuliana.Deposit(1000)
	accountJuliana.ToWithdraw(200)
	PayBankingBillet(&accountJuliana, 300)
	fmt.Println(accountJuliana)
}
