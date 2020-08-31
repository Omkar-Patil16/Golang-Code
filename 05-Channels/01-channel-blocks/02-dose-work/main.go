package main

import "fmt"

func main(){
	ch := make(chan string)

	go func(){
			ch <- "Naruto"

	}()
	fmt.Println(<-ch)

}
