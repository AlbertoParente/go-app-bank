package main

import (
	"fmt"
	"github.com/go-app-bank/accounts"
	"github.com/go-app-bank/customers"
)

func main() {

	clientAlberto := customers.Holder{
        Name: "Alberto",
        CPF: "123.123.123.12",
        Profession: "Dev"}

	accountAlberto := accounts.CurrentAccount{clientAlberto, 123, 123456}
	fmt.Println(accountAlberto)
	
	accountExample := accounts.CurrentAccount{}
    accountExample.Deposit(100)

    fmt.Println(accountExample.GetBalance())

}
