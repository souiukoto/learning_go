package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type square struct {
	length float64
	width  float64
}

type shape interface {
	area() float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (s square) area() float64 {
	return s.length * s.width
}

func info(s shape) {
	fmt.Println("INFO:\n\t", s.area())
}

func main() {
	c := circle{
		radius: 5,
	}

	s := square{
		length: 10,
		width:  5,
	}

	fmt.Println("CIRCLE AREA:\n\t", c.area())
	info(c)
	fmt.Println("SQUARE AREA:\n\t", s.area())
	info(s)
}
