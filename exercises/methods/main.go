package main

import (
	"fmt"

	"github.com/haeju/Go-study/exercises/methods/accounts"
)

func main() {
	account := accounts.NewAccount("james")
	account.Deposit(10)
	fmt.Println(account)
}
