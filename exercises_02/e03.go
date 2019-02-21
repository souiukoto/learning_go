package main

import "fmt"

func main() {
	const typed int8 = 4
	const untyped = 2

	fmt.Printf("%T\n", typed)
	fmt.Printf("%T\n", untyped)
}
