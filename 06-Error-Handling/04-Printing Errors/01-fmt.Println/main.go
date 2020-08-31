package main

import (
	"fmt"
	"os"
)

func main(){

	_,err:=os.Open("rawFile.txt")
	if err!=nil{
		fmt.Println("THERE IS AN ERROR",err)

	}

}
