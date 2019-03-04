package main

import "fmt"

func main() {
	currentYear := 1991

	for currentYear < 2020 {
		fmt.Println(currentYear)
		currentYear += 1
	}
}
