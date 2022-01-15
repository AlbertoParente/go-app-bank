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
	
}
