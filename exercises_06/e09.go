package main

import "fmt"

func sum(nums ...int) int {
	total := 0

	for _, v := range nums {
		total += v
	}

	return total
}

func forEach(arr []int, cb func(nums ...int) int) int {
	return cb(arr...)
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(forEach(arr, sum))
}
