package main

import "fmt"

func main(){
	ch := make(chan string)


	ch <- "Naruto"
	fmt.Println(<-ch)

}
