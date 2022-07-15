package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	const totit = 100
	var counter int = totit
	var wg sync.WaitGroup
	var mux sync.Mutex

	wg.Add(totit)
	for i := 0; i < totit; i++ {
		go func() {
			mux.Lock()
			var tmp int = counter
			runtime.Gosched()
			tmp--
			counter = tmp
			mux.Unlock()
			wg.Done()
		}()
		fmt.Println("Number of goroutines: ", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Counter: ", counter)
}
