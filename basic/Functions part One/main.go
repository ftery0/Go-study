package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	totalLenght, _ := lenAndUpper("jerry")
	fmt.Println(totalLenght)
	repeatMe("jerry", "james", "jane")
}