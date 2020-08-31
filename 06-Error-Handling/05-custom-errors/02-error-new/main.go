package main

import (
	"errors"
	"fmt"
	"log"
)

var custer = errors.New("Sqrt of negative number")


func sqr(f float64) (float64, error) {
	fmt.Printf("%T\n",custer)
	if f < 0 {
		return f,custer
	}
	return 23, nil
}

func main() {
	_, err := sqr(-10)
	if err != nil {
		log.Fatalln(err)
	}

}
