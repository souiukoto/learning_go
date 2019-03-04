package main

import "fmt"

func main() {
	currentYear := 1991

	for {
		fmt.Println(currentYear)
		currentYear += 1

		if currentYear == 2020 {
			break
		}
	}
}
