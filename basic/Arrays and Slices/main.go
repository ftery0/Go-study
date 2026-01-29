package main

import "fmt"

func main() {
	names := []string{"jenny", "james", "dale"}
	names = append(names, "jule")
	fmt.Println(names)

}