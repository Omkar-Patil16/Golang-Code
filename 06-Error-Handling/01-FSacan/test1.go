package main

import "fmt"

func main(){


	var inp, inp2 string
	fmt.Println("Name:")
	_,err:=fmt.Scan(&inp)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println("Last:")
	_,err=fmt.Scan(&inp2)
	if err!=nil{
		fmt.Println(err)
	}


}
