package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = false

func main() {
	fmt.Println("EXERCISE 3:")
	fmt.Println("===========")
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(s)
	fmt.Println("===========")
}
