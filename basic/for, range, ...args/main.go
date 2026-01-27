package main

import "fmt"

func canIDrink(age int) bool {
	for i:=0; i<age; i++ {
		fmt.Println(i)
	}
	return true
}


func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	result := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(result)
	canIDrink(18)
}