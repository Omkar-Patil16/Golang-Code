package main

import "fmt"



type person struct{
	name string
	age int

}

type human interface {

	speak()
}

func (p *person) speak() {
	fmt.Println("My Name is ",p.name," and i am ",p.age," years old ")
}

func saySomething(t human){
	t.speak()
}


func main(){

	p1 :=&person{"Naruto",17}
	saySomething(p1)

}
