package main

import (
	"fmt"
	"sync"
)

var cot int

func main() {
	gr := 10
	var wg sync.WaitGroup
	var mux sync.Mutex
	wg.Add(gr)

	for i := 0; i < gr; i++ {
		go func() {
			mux.Lock()
			ct := cot
			ct++
			cot = ct
			fmt.Println( cot)
			mux.Unlock()
			wg.Done()

		}()

	}

	fmt.Println("Count", cot)
	wg.Wait()
}
