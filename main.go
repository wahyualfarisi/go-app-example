package main

import (
	"fmt"

	gosayhello "github.com/wahyualfarisi/go-say-hello"
)

func main() {
	// say hello function
	sh := gosayhello.SayHello("Wahyu")
	fmt.Printf("This is a value of %s\n", sh)

	// calculate sum
	result := gosayhello.CalculateSum([]int{1, 2, 3})
	fmt.Println(result)
}
