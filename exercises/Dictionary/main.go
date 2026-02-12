package main

import (
	"fmt"

	"github.com/haeju/Go-study/exercises/Dictionary/dic"
)

func main() {
	dictionary := dic.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	dictionary.Search(baseWord)
	dictionary.Delete(baseWord)
	word, err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(word)
	}
}