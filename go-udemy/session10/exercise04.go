package main

import "fmt"

// comma, ok - & - select
func main() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	v, ok := <-c
	fmt.Println(v, ok)
	close(c)

	v, ok = <-c
	fmt.Println(v, ok)

	// _______

	co := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			co <- i
		}
		close(co)
	}()

	for v := range co {
		fmt.Println(v)
	}

	fmt.Println("Finishing")
}
