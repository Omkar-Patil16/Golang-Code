package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)



func main() {
	var at int64

	gr := 10
	var wg sync.WaitGroup
	wg.Add(gr)

	for i := 0; i < gr; i++ {
		go func() {
			atomic.AddInt64(&at,4)
			fmt.Println( "Counter",atomic.LoadInt64(&at))
			wg.Done()

		}()

	}


	wg.Wait()
	fmt.Println("End Value",at)
}
