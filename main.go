package main

import "fmt"

type CurrentAccount struct{
	holder string
	agencyNumber int
	accountNumber int
	balance float64
}

func main() {
	accountAlberto := CurrentAccount{holder: "Alberto Parente", agencyNumber: 123, accountNumber: 123456, balance: 150.00}
	accountJuliana := CurrentAccount{"Alberto Parente", 321, 654321, 100.00}

	fmt.Println(accountAlberto, accountJuliana)
}