package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}

	fmt.Println("===========")
	fmt.Println("Array range:")
	for i, v := range arr {
		fmt.Println(i, v)
	}
	fmt.Println("===========")
	fmt.Printf("Type: %T\n", arr)
	fmt.Println("===========")
}
