package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// divide a job into several goroutines
func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go agregar(c1)

	go fanOutIn(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}

	fmt.Println("Finishing...")
}

func agregar(c chan<- int) {
	for i := 0; i < 20; i++ {
		c <- i
	}
	close(c)
}

func fanOutIn(c1, c2 chan int) {
	var wg sync.WaitGroup
	const ngorutinas = 10
	wg.Add(ngorutinas)

	for i := 0; i < ngorutinas; i++ {
		go func() {
			for v := range c1 {
				func(v2 int) {
					c2 <- trabajoConsumeTiempo(v2)
				}(v)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	close(c2)
}

func trabajoConsumeTiempo(n int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(1000)
}
