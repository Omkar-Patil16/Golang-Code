package main

import (
	"fmt"
	"github.com/Omkar-Patil16/Golang-Code/10-Ninja-Level-12/dog"
)

type  bulldog struct {
	name string
	age int
}

func main(){

		bd1 := bulldog{
			name:"root",
			age : dog.Years(5),
		}
		fmt.Println(bd1)
}
