package main

import "fmt"

type customErr struct{
	info string
}

func (ce customErr) Error() string{
	return  fmt.Sprintf("this is from Error() --> %v",ce.info)
}

func foo(e error){
	fmt.Println("From foo () --->",e," \n ",e.(customErr).info)

}
func main(){
	ce1 := customErr{
		info :" From customErr need a drink ",
	}


	foo(ce1)
}