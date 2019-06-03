package main

import "fmt"

func main() {
	slice := make([]int, 10, 100)
	nums := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	for i, v := range nums {
		slice[i] = v
	}

	fmt.Println("=======================")
	sliceA := slice[:5]
	sliceB := slice[5:]
	sliceC := slice[1:6]
	fmt.Println(sliceA, "sliceA")
	fmt.Println(sliceB, "sliceB")
	fmt.Println(sliceC, "sliceC")
	fmt.Println("=======================")
}
