package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	c := make(chan int)

	wg.Add(1)
	go func() {
		c <- 42
		wg.Done()
	}()
	fmt.Println("Routines: ", runtime.NumGoroutine())
	fmt.Println(<-c)
	wg.Wait()
	fmt.Println("Routines: ", runtime.NumGoroutine())
}
