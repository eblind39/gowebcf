package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer foos()

	_, err := os.Open("no-file.txt")
	if err != nil {
		log.Panicln(err)
		// panic(err) // same as above
	}
}

func foos() {
	fmt.Println("foos executes after a call to panic")
}
