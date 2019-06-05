package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("sumVariadic:\n\t", sumVariadic(arr...))
	fmt.Println("sumFromArray:\n\t", sumFromArray(arr))
}

func sumVariadic(nums ...int) int {
	s := 0

	fmt.Println("nums:\n\t", nums)

	for _, v := range nums {
		s += v
	}

	return s
}

func sumFromArray(nums []int) int {
	s := 0

	fmt.Println("nums:\n\t", nums)

	for _, v := range nums {
		s += v
	}

	return s
}
