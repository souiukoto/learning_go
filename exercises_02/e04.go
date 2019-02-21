package main

import "fmt"

func main() {
	num := 42
	fmt.Printf("%d\n", num)
	fmt.Printf("%b\n", num)
	fmt.Printf("%X\n", num)
	shiftedNum := 42 << 1
	fmt.Printf("%d\n", shiftedNum)
	fmt.Printf("%b\n", shiftedNum)
	fmt.Printf("%X\n", shiftedNum)
}
