package main

import "fmt"

type person struct {
	first string
	last  string
}

func changeMe(p *person) *person {
	*p = person{
		first: "Bob",
		last:  "Brown",
	}

	return p
}

func main() {
	p := person{
		first: "Peter",
		last:  "McKinley",
	}

	fmt.Println("Original entry at address:")
	fmt.Println("\t", *&p)
	fmt.Println("New entry at address:")
	fmt.Println("\t", *changeMe(&p))
}
