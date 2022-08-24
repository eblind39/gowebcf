package main

import (
	"github.com/fatih/color"
	"time"
)

func server1(ch chan string) {
	for {
		time.Sleep(6 * time.Second)
		ch <- "This is from server 1"
	}
}

func server2(ch chan string) {
	for {
		time.Sleep(3 * time.Second)
		ch <- "This is from server 2"
	}
}

func main() {
	color.Cyan("Select with channels")
	color.Cyan("--------------------")

	channel1 := make(chan string)
	channel2 := make(chan string)

	go server1(channel1)
	go server2(channel2)

	for {
		select {
		case s1 := <-channel1:
			color.Blue("Case one: %s", s1)
		case s2 := <-channel1:
			color.Blue("Case two: %s", s2)
		case s3 := <-channel2:
			color.Green("Case three: %s", s3)
		case s4 := <-channel2:
			color.Green("Case four: %s", s4)
			// default:
			// avoiding deadlock
		}
	}
}
