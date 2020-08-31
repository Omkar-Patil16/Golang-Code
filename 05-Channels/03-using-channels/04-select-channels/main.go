package main

import "fmt"

func sendTo(e,o,q chan<- int){
	for i:=0;i<10;i++{
		if i%2==0{
			e<-i
		}else{
			o<- i
		}

	}
	q<-0

}
func main(){

	even :=  make(chan int)
	odd :=  make(chan int)
	quit :=  make(chan int)

	go sendTo(even,odd,quit)

	for {
		select {
			case v := <-even :
				fmt.Println("even",v)
			case v := <-odd :
				fmt.Println("odd ",v)
			case v := <-quit :
				fmt.Println("exit",v)
			return
		}
	}



}

