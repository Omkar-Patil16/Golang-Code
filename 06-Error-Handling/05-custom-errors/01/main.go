package main

import (
	"errors"
	"log"
)

func sqr(f float64) (float64, error) {
	if f < 0 {
		return f, errors.New("Sqrt of negative number")
	}
	return 23, nil
}

func main() {
	_, err := sqr(-10)
	if err != nil {
		log.Fatalln(err)
	}

}
