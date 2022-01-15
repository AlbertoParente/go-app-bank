package main

import (
	"fmt"

	"github.com/go-app-bank/accounts"
)

func main() {

	accountAlberto := accounts.CurrentAccount{
		Holder:        "Alberto Parente",
		AgencyNumber:  987,
		AccountNumber: 987654321,
		Balance:       1000,
	}

	accountJuliana := accounts.CurrentAccount{
		Holder:        "Juliana",
		AgencyNumber:  987,
		AccountNumber: 987654321,
		Balance:       1000,
	}

	status := accountAlberto.Transfer(-200, &accountJuliana)

	fmt.Println(status)
	fmt.Println(accountAlberto)
	fmt.Println(accountJuliana)

	/*
		accountAlberto := accounts.CurrentAccount{}
		accountAlberto.Holder = "Alberto Parente"
		accountAlberto.AgencyNumber = 123
		accountAlberto.AccountNumber = 123456789
		accountAlberto.Balance = 500

		fmt.Println(accountAlberto.Balance)
		fmt.Println(accountAlberto.ToWithdraw(300))
		fmt.Println(accountAlberto.Balance)
		fmt.Println(accountAlberto.ToWithdraw(500))
		fmt.Println(accountAlberto.Balance)
		fmt.Println(accountAlberto.ToWithdraw(-300))
		fmt.Println(accountAlberto.Balance)

		fmt.Println(accountAlberto.Balance)
		status, value := accountAlberto.Deposit(2000)
		fmt.Println(status, value)

		accountJuliana := accounts.CurrentAccount{
			Holder:        "Juliana",
			AgencyNumber:  987,
			AccountNumber: 987654321,
			Balance:       100,
		}

		status2 := accountAlberto.Transfer(200, &accountJuliana)

		fmt.Println(status2)
		fmt.Println(accountJuliana)
		fmt.Println(accountAlberto)
	*/
}
