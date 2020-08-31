package main

import "fmt"

func main(){
	ch := make(chan <-int,1)
	fmt.Printf("%T",ch)
	fmt.Println(<-ch)
}
