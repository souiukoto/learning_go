package main

import "fmt"

func main() {
	if true && false {
		fmt.Println("Not entered")
	} else if false || false {
		fmt.Println("Not entered")
	} else {
		fmt.Println("Entered")
	}
}
