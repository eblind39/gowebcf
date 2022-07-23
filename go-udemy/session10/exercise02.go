package main

import (
	"fmt"
)

func main() {

	c := gen() // receive only channel
	recibirr(c)

	fmt.Println("Finishing...")
}

func gen() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 40; i++ {
			c <- i
		}
		// We will use range, so close the channel
		close(c)
	}()
	return c
}

func recibirr(c <-chan int) {
	// always close the chan when We use range
	for v := range c {
		fmt.Println(v)
	}

}
