package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	f, err := os.Create("newFile.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	r := strings.NewReader("microsoft flight simulator")
	_, err = io.Copy(f, r)
	if err != nil {
		log.Fatal(err)
	}

}
