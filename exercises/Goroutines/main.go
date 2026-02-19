package main

import (
	"fmt"
	"time"
)

func main() {
	go count("Alice")
	go count("Bob")

	time.Sleep(5 * time.Second)
}

func count(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, ":", i)
		time.Sleep(time.Second)
	}
}