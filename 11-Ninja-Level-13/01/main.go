package main

import (
	"fmt"
	Dog "github.com/Omkar-Patil16/Golang-Code/11-Ninja-Level-13/01/Dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  Dog.Years(10),
	}
	fmt.Println(fido)
	fmt.Println(Dog.YearsTwo(20))
}
