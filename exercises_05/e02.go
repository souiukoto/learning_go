package main

import "fmt"

type person struct {
	first                   string
	last                    string
	favoriteIceCreamFlavors []string
}

func main() {
	p1 := person{
		first: "Peter",
		last:  "McKinley",
		favoriteIceCreamFlavors: []string{
			"Banana",
			"Pistachio",
			"Cookies & Cream",
		},
	}

	p2 := person{
		first: "Bob",
		last:  "Bobbington",
		favoriteIceCreamFlavors: []string{
			"Nocciola",
		},
	}

	people := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	fmt.Println(people)
}
