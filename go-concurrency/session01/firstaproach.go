package main

import (
	"fmt"
	"sync"
)

func printSomething(s string) {
	fmt.Println(s)
}

func main() {
	words := []string{
		"alpha",
		"beta",
		"delta",
		"gamma",
		"pi",
		"zeta",
		"eta",
		"tetha",
		"epsilon",
	}

	var wg sync.WaitGroup

	wg.Add(len(words))

	for i, x := range words {
		func() {
			go printSomething(fmt.Sprintf("%d: %s", i, x))
			wg.Done()
		}()
	}
	wg.Wait()

	printSomething("This is the last thing to be printed!")
}
