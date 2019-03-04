package main

import "fmt"

func main() {
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	for _, letter := range alphabet {
		fmt.Printf("%#U\n", letter)
		fmt.Printf("%#U\n", letter)
		fmt.Printf("%#U\n", letter)
	}
}
