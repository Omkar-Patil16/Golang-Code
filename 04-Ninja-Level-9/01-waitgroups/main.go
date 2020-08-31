package main

import (
"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	fmt.Println("Goroutine",runtime.NumGoroutine())

	fmt.Println("Hello, playground")
	wg.Add(2)

	go first()

	go Second()
	fmt.Println("Goroutine",runtime.NumGoroutine())


	wg.Wait()

}

func first() {
	fmt.Println("First Routine")
	wg.Done()

}
func Second() {
	fmt.Println("Second Routine")
	wg.Done()
}

