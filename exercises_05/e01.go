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

	fmt.Println(p1.first, p1.last)
	for _, v := range p1.favoriteIceCreamFlavors {
		fmt.Println("\t", v)
	}
	fmt.Println("==================")
	fmt.Println(p2.first, p2.last)
	for _, v := range p2.favoriteIceCreamFlavors {
		fmt.Println("\t", v)
	}
}
