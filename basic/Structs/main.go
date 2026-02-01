package main

import "fmt"

type Person struct {
	Name    string
	Age     int
	Address string
	Hobby   []string
}


func main() {

	person1 := Person{
		Name:    "haejun",
		Age:     20,
		Address: "seoul",
		Hobby:   []string{"football", "programming"},
	}
	fmt.Println(person1.Name)	
	fmt.Println(person1.Age)
	fmt.Println(person1.Address)
	fmt.Println(person1.Hobby)
}