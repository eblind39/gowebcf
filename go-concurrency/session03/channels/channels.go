package main

import (
	"fmt"
	"github.com/fatih/color"
	"strings"
)

func shout(ping <-chan string, pong chan<- string) {
	for {
		s := <-ping

		pong <- fmt.Sprintf("%s!!!", strings.ToUpper(s))
	}
}

func main() {
	// create two channels
	ping := make(chan string)
	pong := make(chan string)

	go shout(ping, pong)

	color.Cyan("Type something and press ENTER (enter Q to quit)")

	for {
		// print a prompt
		color.Cyan("-> ")

		// get user input
		var userInput string
		_, _ = fmt.Scanln(&userInput)

		if userInput == strings.ToLower("q") {
			break
		}

		ping <- userInput
		// wait for a response

		response := <-pong
		color.Cyan("Response: %s\n", response)
	}

	color.Green("All done. Closing channels.")
	close(ping)
	close(pong)
}
