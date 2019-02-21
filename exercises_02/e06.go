package main

import "fmt"

func main() {
	const currentYear int = 2019

	const (
		first  int = currentYear + iota
		second int = currentYear + iota
		third  int = currentYear + iota
		fourth int = currentYear + iota
	)

	fmt.Println(first)
	fmt.Println(second)
	fmt.Println(third)
	fmt.Println(fourth)
}
