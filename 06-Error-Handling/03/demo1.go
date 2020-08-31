package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main(){

	f,err := os.Open("newFile.txt")
	if err!=nil{
		log.Fatal(err)
	}
	defer f.Close()

	bs,err := ioutil.ReadAll(f)
	if err!=nil{
		log.Fatal(err)
	}

	fmt.Println(string(bs))

}
