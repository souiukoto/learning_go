package main

import "fmt"

type vehicle struct {
	color string
	doors int
}

type sedan struct {
	vehicle
	luxury bool
}

type truck struct {
	vehicle
	fourWheel bool
}

func main() {
	t1 := truck{
		vehicle: vehicle{
			color: "red",
			doors: 2,
		},
		fourWheel: true,
	}

	s1 := sedan{
		vehicle: vehicle{
			color: "blue",
			doors: 4,
		},
		luxury: true,
	}

	fmt.Println(s1)
	fmt.Println(t1)
}
