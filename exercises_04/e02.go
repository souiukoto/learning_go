package main

import "fmt"

func main() {
	slice := make([]int, 10, 100)

	for i := range slice {
		slice[i] = i + 1
	}

	fmt.Println("===========")
	fmt.Println("Slice range:")
	for i, v := range slice {
		fmt.Println(i, v)
	}
	fmt.Println("===========")
	fmt.Printf("Type: %T\n", slice)
	fmt.Println("===========")
}
