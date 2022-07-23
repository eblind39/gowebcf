package main

import (
	"fmt"
)

func main() {
	salir := make(chan int)
	c := gens(salir)

	recibirs(c, salir)

	fmt.Println("Finishing...")
}

func gens(salir chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		salir <- 1
		close(c)
	}()

	return c
}

func recibirs(c, salir <-chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-salir:
			fmt.Println("No more values...")
			return
		}
	}
}
