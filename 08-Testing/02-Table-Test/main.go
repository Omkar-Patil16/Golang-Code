package main

import "fmt"

func mySub(x ...int) int {
	sub := 0
	for _,v := range x{

		sub = v - sub
		if sub < 0 {
			sub = - sub
		}
	}
	return sub + 1
}


func main(){
	fmt.Println("9-4 is = ", mySub(9,4,1,3))
	fmt.Println("8-5 is = ", mySub(8,5))
	fmt.Println("7-1 is = ", mySub(7,1))

}