package main

import (
	"fmt"
)

func main() {

	var text1, text2 string
	fmt.Print("Name:")
	_, err := fmt.Scan(&text1)
	if err != nil {
		panic(err)
	}

	fmt.Print("Name:")
	_, err = fmt.Scan(&text2)
	if err != nil {
		panic(err)
	}

}

