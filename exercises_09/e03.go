package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	x := 0
	routines := 10
	wg.Add(routines)

	for i := 0; i < routines; i++ {
		go func() {
			v := x
			v++
			x = v
			fmt.Println(x)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(x)
}
