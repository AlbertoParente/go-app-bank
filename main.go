package main

import (
	"fmt"

	"github.com/go-app-bank/accounts"
)

func main() {

	ClientAlberto := accounts.CurrentAccount{Holder: customers.Holder{
        Name: "Alberto",
        CPF: "123.123.123.12",
        Profession: "Dev"},
        AgencyNumber:123, 
		AccountNumber: 123456, 
		Balance:100
	}

	accountAlberto := accounts.CurrentAccount(
		ClientAlberto, 
		123, 
		123456, 
		100
		fmt.Println(accountAlberto)
	
	accountExample := accounts.CurrentAccount{}
    accountExample.Deposit(100)

    fmt.Println(accountExample.GetBalance())

}
