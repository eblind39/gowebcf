package main

import "fmt"

func main() {
	// unbuffered channel
	// we need an extra goroutine to send a value
	// then, the main goroutine receives its value
	c1 := make(chan int32)

	go func() {
		c1 <- 42
		c1 <- 43
	}()

	fmt.Println(<-c1)
	fmt.Println(<-c1)

	// buffered channel
	// we can set and get the value of the channel
	// in the same (main) goroutine, we say to Go
	// the quantity of values the channel can manage
	c2 := make(chan int32, 2) // if value==1, we have a deadlock

	c2 <- 70
	c2 <- 71

	fmt.Println(<-c2)
	fmt.Println(<-c2)

	// print the type of channel
	fmt.Printf("%T\n", c2)

	// "sent only" channel
	cs := make(chan<- int32, 1)
	cs <- 91
	// fmt.Println(<-c3) - // error, not receiving channel
	fmt.Printf("%T\n", cs)

	// "receive only" channel
	cr := make(<-chan int32, 1)
	fmt.Printf("%T\n", cr)

	// cast bidirectional channel to a specific direction
	fmt.Printf("%T\n", (<-chan int32)(c2))
	fmt.Printf("%T\n", (chan<- int32)(c2))

	// we can assign (inferred type) a specific (direction) channel
	// from a bidirectional channel, Go determines which direction
	// resolve into the assignation
	cr = c2
	cs = c1
}
