package main

import (
	"fmt"
	"sync"
)

// Fan In - Design pattern for concurrent programming
func main() {
	odd := make(chan int)
	even := make(chan int)
	fanin := make(chan int)

	// send
	go sendvalss(odd, even)

	// receive
	go receivevalss(odd, even, fanin)

	for v := range fanin {
		fmt.Println(v)
	}

	fmt.Println("Finishing...")
}

func sendvalss(odd, even chan<- int) {
	for i := 0; i < 40; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(odd)
	close(even)
}

func receivevalss(odd, even <-chan int, fanin chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range odd {
			fanin <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range even {
			fanin <- v
		}
		wg.Done()
	}()

	wg.Wait()
	close(fanin)
}
