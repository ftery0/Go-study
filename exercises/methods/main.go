package main

import (
	"fmt"

	"github.com/haeju/Go-study/exercises/methods/accounts"
)

func main() {
	account := accounts.NewAccount("james")
	account.Deposit(10)
	fmt.Println(account.Balance())


	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account.Balance())

}
