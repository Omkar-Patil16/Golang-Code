package main

import "fmt"

func main(){
	// create channel

	ch := make(chan int)
	//send channel
	go ping(ch)
	//receive channel
	go drop(ch)

}

//send func
func ping(cs chan<- int){
	cs <- 24
}

//receive func
func drop(cr <-chan int){
	fmt.Println{"Channel",<-cr}
}