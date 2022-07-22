package main

import "fmt"

// comma, ok: ensures to read only valid values from a ready-to-receive channel
func main() {
	odd := make(chan int)
	even := make(chan int)
	exit := make(chan bool)

	// send
	go sendvals(odd, even, exit)

	// receive
	receivevals(odd, even, exit)

	fmt.Println("Finishing...")

	c := make(chan int)
	go func() {
		c <- 42
		close(c)
	}()
	v, ok := <-c
	fmt.Println(v, ok)
	v, ok = <-c
	fmt.Println(v, ok)
}

func sendvals(odd, even chan<- int, exit chan<- bool) {
	for i := 0; i < 30; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(exit)
}

func receivevals(odd, even <-chan int, exit <-chan bool) {
	for {
		select {
		case v := <-odd:
			fmt.Println("From channel 'odd': ", v)
		case v := <-even:
			fmt.Println("From channel 'even': ", v)
		case i, ok := <-exit:
			if !ok {
				fmt.Println("From comma ok: ", i, ok)
				return
			} else {
				fmt.Println("From comma ok: ", i)
			}
			return
		}
	}
}
