package main

import "fmt"

func main() {
	fmt.Println("Calling foo:")
	fmt.Println("\t", foo())
	fmt.Println("Calling bar:")
	x, y := bar()
	fmt.Println("\t", x, y)
}

func foo() int {
	return 10
}

func bar() (int, string) {
	return 20, "Bucks"
}
