package main



import "fmt"

func main() {
	mapstudy := map[string]string{"name": "haejun", "age": "20", "address": "seoul"}
	for key, value := range mapstudy {
		fmt.Println(key, value)
	}
}