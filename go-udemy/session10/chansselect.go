package main

import "fmt"

// select: retrieves values from channels ready to be extracted
func main() {
	odd := make(chan int)
	even := make(chan int)
	exit := make(chan int)

	// send
	go sendval(odd, even, exit)

	// receive
	receiveval(odd, even, exit)

	fmt.Println("Finishing...")
}

func sendval(odd, even, exit chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	// close(odd)
	// close(even)
	exit <- 0
}

func receiveval(odd, even, exit <-chan int) {
	for {
		select {
		case v := <-odd:
			fmt.Println("From channel 'odd': ", v)
		case v := <-even:
			fmt.Println("From channel 'even': ", v)
		case v := <-exit:
			fmt.Println("From channel 'exit': ", v)
			return
		}
	}
}
