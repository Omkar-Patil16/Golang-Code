package main

import (
	"fmt"
	"github.com/Omkar-Patil16/Golang-Code/11-Ninja-Level-13/02/quote"
	"github.com/Omkar-Patil16/Golang-Code/11-Ninja-Level-13/02/word"
)

func main() {


	for k, v := range word.UseCount(quote.SunAlso) {


		fmt.Println(v," : ", k)
	}

	fmt.Println("\nTotal of Word in Quote :",word.Count(quote.SunAlso))
}
