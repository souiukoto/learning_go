package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Foo:")
	wg.Add(3)
	go foo()
	fmt.Println("===========================")
	fmt.Println("FizzBuzz:")
	go fizzBuzz()
	go bar()
	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	wg.Done()
}

func bar() {
	fmt.Println("BOO!")
	wg.Done()
}

func fizzBuzz() {
	for i := 1; i < 100; i++ {
		if i%15 == 0 {
			fmt.Println(i, "fizz buzz")
		} else if i%5 == 0 {
			fmt.Println(i, "buzz")
		} else if i%3 == 0 {
			fmt.Println(i, "fizz")
		}
	}

	wg.Done()
}
