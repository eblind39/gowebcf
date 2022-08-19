package main

import (
	"fmt"
)

var msg string

func updateMessage(s string) {
	msg = s
}

func printMessage() {
	fmt.Println(msg)
}

func main() {
	// challenge: modify this code so that the calls to updateMessage() on line
	// 27, 30, and 33 run as goroutines, and implement WaitGroup so that
	// the program runs properly, and prints out three different messages.
	// Then, write a test for all three functions in this program: updateMessage(),
	// printMessage, and main()

	msg = "Hello, world!"

	updateMessage("Hello, universe!")
	printMessage()

	updateMessage("Hello, cosmos!")
	printMessage()

	updateMessage("Hello, world!")
	printMessage()
}
