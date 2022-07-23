package main

import (
	"fmt"
	"sync"
)

func main() {
	co := make(chan int)
	var wg sync.WaitGroup

	wg.Add(10)

	for j := 0; j < 10; j++ {
		go func() {
			for i := 0; i < 1; i++ {
				co <- i
			}
			wg.Done()
		}()
	}

	for k := 0; k < 10; k++ {
		fmt.Println(<-co)
	}

	fmt.Println("Finishing")
	wg.Wait()
}
