package main

import "fmt"

func sendToCh(send chan<- string,gt string){
	send <- gt
}

func recFromCh(send <-chan string, recv chan<- string){
	msg := <-send
	recv<- msg
}

func main(){
			send := make(chan string)
			recv := make(chan string)

			go sendToCh(send,"message")
			go recFromCh(send,recv)
			fmt.Println(<-recv)

}
