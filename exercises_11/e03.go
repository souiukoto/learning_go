package main

import "fmt"

type customError struct {
	info string
}

func (ce customError) Error() string {
	return fmt.Sprintf("Error: %v", ce.info)
}

func main() {
	c1 := customError{
		info: "Custom Error",
	}

	foo(c1)
}

func foo(e error) {
	fmt.Println("FOO", e)
}
