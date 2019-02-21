package main

import "fmt"

func main() {
	isEqual := 42 == 42
	isLessOrEqual := 42 <= 42
	isGreaterOrEqual := 42 >= 43
	doesNotEqual := 42 != 42
	lessThan := 42 < 42
	greaterThan := 42 > 42

	fmt.Println(isEqual)
	fmt.Println(isLessOrEqual)
	fmt.Println(isGreaterOrEqual)
	fmt.Println(doesNotEqual)
	fmt.Println(lessThan)
	fmt.Println(greaterThan)
}
