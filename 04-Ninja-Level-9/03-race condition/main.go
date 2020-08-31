package main

import (
	"fmt"
	"runtime"
	"sync"
)

var cot int

const gr = 100

var wg sync.WaitGroup

func main() {
	wg.Add(gr)
	for i := 0; i < gr; i++ {
		go func() {
			ct := cot
			runtime.Gosched()
			ct++
			cot = ct
			wg.Done()

		}()
		fmt.Println("Count", cot)
	}
	fmt.Println("GoRot", runtime.NumGoroutine())
	fmt.Println("Count", cot)
	wg.Wait()
}
