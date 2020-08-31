package main

import (
	"fmt"
	"log"
	"os"
)

func myPrint() {
	fmt.Println("Testing defer function")
}
func main() {
	defer myPrint()
	_, err := os.Open("rawFile.txt")
	if err != nil {
		log.Panicln("THERE IS AN ERROR", err)

	}

}
