package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func main() {
	var x int64
	routines := 10
	wg.Add(routines)

	for i := 0; i < routines; i++ {
		go func() {
			atomic.AddInt64(&x, 1)
			fmt.Println(atomic.LoadInt64(&x))
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(x)
}
