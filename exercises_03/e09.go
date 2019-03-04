package main

import "fmt"

func main() {
	favSport := "Swimming"

	switch favSport {
	case "Swimming":
		fmt.Println("Entered Swimming")
	case "Boxing":
		fmt.Println("Not entered")
	}
}
