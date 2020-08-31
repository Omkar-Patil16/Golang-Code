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
	return sub + 2
}


func main(){
	fmt.Println("9-4 is = ", mySub(9,4))
	fmt.Println("8-2 is = ", mySub(8,2))
	fmt.Println("6-1 is = ", mySub(6,1))

}