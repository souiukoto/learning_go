package main

import "fmt"

func main() {
	fmt.Println("Variable func:")
	f := func(x int, y int) int {
		return x * y
	}

	fmt.Println("\t", f(10, 5))
}
