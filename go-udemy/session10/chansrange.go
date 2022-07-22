package main

import (
	"fmt"
)

// range reads items from a channel whilst this channel is opened
func main() {
	c := make(chan int)

	// send
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()

	// receive
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("Finishing...")
}
