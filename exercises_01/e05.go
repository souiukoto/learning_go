package main

import "fmt"

type customType int

var customInstance customType
var intInstance int

func main() {
	fmt.Println("EXERCISE 5:")
	fmt.Println("===========")
	fmt.Println(customInstance)
	fmt.Printf("%T\n", customInstance)
	customInstance = 10
	fmt.Println(customInstance)
	intInstance = int(customInstance)
	fmt.Println(intInstance)
	fmt.Printf("%T\n", intInstance)
	fmt.Println("===========")
}
