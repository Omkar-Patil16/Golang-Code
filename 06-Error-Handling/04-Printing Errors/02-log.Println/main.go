package main

import (
	"log"
	"os"
)

func main(){

	_,err:=os.Open("rawFile.txt")
	if err!=nil{
		log.Println("THERE IS AN ERROR",err)

	}

}
