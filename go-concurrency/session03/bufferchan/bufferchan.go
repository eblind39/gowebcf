package main

import (
	"github.com/fatih/color"
	"time"
)

func listenToChan(ch chan int) {
	for {
		// print a got data message
		i := <-ch
		color.Cyan("Got [%d] from channel", i)

		// simulate doing a lot of work
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	// allow up to 10 int's in the channel at the same time
	ch := make(chan int, 10)

	go listenToChan(ch)

	for i := 0; i <= 100; i++ {
		color.Cyan("Sending [%d] to channel...", i)
		ch <- i
		color.Cyan("Sent [%d] to channel!", i)
	}

	color.Green("Done!")
	close(ch)
}
