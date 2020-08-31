package main

import "fmt"

func sendTo(so chan<- int) {
	for i := 0; i < 10; i++ {
		so <- i
	}
	close(so)
}

func recvFrom(rev <-chan int) {
	for v := range rev {

		fmt.Println(v)
	}
}

func main() {
	ch := make(chan int)

	go sendTo(ch)
	 recvFrom(ch)

}
