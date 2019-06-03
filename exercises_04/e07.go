package main

import "fmt"

func main() {
	bond := []string{"James", "Bond", "Shaken, not stirred"}
	mp := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	arr := [][]string{bond, mp}
	fmt.Println("==============================")

	for _, v := range arr {
		fmt.Println(v)

		for _, w := range v {
			fmt.Println(w)
		}
	}
}
