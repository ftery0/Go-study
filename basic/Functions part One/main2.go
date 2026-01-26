package main

import (
	"fmt"
	"strings"
)


func lenAndUpper2(name string) (lenght int, uppercase string) {
	defer fmt.Println("I'm done") //defer is used to delay the execution of a function until the surrounding function returns
	lenght = len(name)
	uppercase = strings.ToUpper(name)
	return // naked return

}

func main2() {
	totalLenght, up := lenAndUpper("jerry")
	fmt.Println(totalLenght, up)
}