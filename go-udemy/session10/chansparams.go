package main

import (
	"fmt"
)

func main() {
	c := make(chan int32)

	// extra goroutine to send a value to the unbuffered channel
	go send(c)
	receive(c)

	fmt.Println("Finishing...")
}

func send(c chan<- int32) {
	c <- 42
}

func receive(c <-chan int32) {
	fmt.Println("Channel value received: ", <-c)
}
