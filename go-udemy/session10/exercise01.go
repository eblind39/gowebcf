package main

import (
	"fmt"
)

func main() {
	// unbuffered channel, needs a goroutine to send a value
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
	close(c)

	// buffered channel
	bc := make(chan int, 1)
	bc <- 43
	fmt.Println(<-bc)
	close(bc)
}
