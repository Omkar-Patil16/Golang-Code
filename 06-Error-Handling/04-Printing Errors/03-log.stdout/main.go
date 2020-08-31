package main

import (
	"log"
	"os"
)

func main(){

	f,err:= os.Create("demo1.txt")
	log.SetOutput(f)
	if err!=nil{
		log.Println("THERE IS AN ERROR",err)
	}
	defer f.Close()

	f2,err:=os.Open("rawFile.txt")
	if err!=nil{
		log.Println("THERE IS AN ERROR",err)

	}
	defer f2.Close()

}
