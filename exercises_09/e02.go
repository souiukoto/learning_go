package main

import "fmt"

func main() {
	p := person{
		first: "Peter",
		last:  "McKinley",
	}

	saySomething(&p)
	// saySomething(p)
}

type person struct {
	first string
	last  string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("Hi, my name is", p.first, p.last)
}

func saySomething(h human) {
	h.speak()
}
