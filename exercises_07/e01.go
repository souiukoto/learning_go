package main

import "fmt"

func main() {
	x := 5
	fmt.Println("Value:")
	fmt.Println("\t", x)
	fmt.Println("Address:")
	fmt.Println("\t", &x)
}
