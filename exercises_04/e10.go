package main

import "fmt"

func main() {
	charMap := map[string][]string{}

	charMap["bond_james"] = []string{
		"Shaken, not stirred",
		"Martinis",
		"Women",
	}

	charMap["moneypenny_miss"] = []string{
		"James Bond",
		"Literature",
		"Computer Science",
	}

	charMap["no_dr"] = []string{
		"Being evil",
		"Ice cream",
		"Sunsets",
	}

	delete(charMap, "no_dr")

	for k, v := range charMap {
		fmt.Println(k)
		for i, val := range v {
			fmt.Println("\t", i, val)
		}
	}
}
