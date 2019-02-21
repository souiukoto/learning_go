package main

import "fmt"

type customType int

var customInstance customType

func main() {
	fmt.Println("EXERCISE 4:")
	fmt.Println("===========")
	fmt.Println(customInstance)
	fmt.Printf("%T\n", customInstance)
	customInstance = 42
	fmt.Println(customInstance)
	fmt.Println("===========")
}
