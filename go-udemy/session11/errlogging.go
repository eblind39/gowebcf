package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer foo()

	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)

	_, err = os.Open("no-file-exists.txt")
	if err != nil {
		fmt.Println("No file existes", err)
		log.Println("An error happened", err)
		log.Fatalln(err) // kills everything, equals to (Println + os.Exist(1))
		panic(err)
	}

	// println formats using the prederteminate formats for its params
	// and writes to the standard output
}

func foo() {
	fmt.Println("foo does not execute after a call to log.Fatal or os.Exit")
}