package main

import (
	"fmt"
	"sync"
)

var m sync.Mutex
var wg sync.WaitGroup

func main() {
	x := 0
	routines := 10
	wg.Add(routines)

	for i := 0; i < routines; i++ {
		go func() {
			m.Lock()
			v := x
			v++
			x = v
			fmt.Println(x)
			m.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(x)
}
