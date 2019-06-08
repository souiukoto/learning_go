package main

import (
	"fmt"
)

func main() {
	cs := make(chan int)
	go send(cs, 42)
	receive(cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}

func send(c chan<- int, num int) {
	c <- num
}

func receive(c <-chan int) {
	fmt.Println(<-c)
}
