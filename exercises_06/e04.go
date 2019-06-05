package main

import "fmt"

type person struct {
	first  string
	gender string
	last   string
}

func (p person) speak() {
	fmt.Println(p.first, p.last)
}

func main() {
	p := person{
		first:  "Peter",
		last:   "McKinley",
		gender: "m",
	}

	p.speak()
}
