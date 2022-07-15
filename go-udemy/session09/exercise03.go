package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		parallelMutex()
		parallelAtomic()
		wg.Done()
	}()
	wg.Wait()
}

func parallelMutex() {
	const totit = 100
	var counter int = totit
	var wg sync.WaitGroup
	var mux sync.Mutex

	wg.Add(totit)
	for i := 0; i < totit; i++ {
		go func() {
			mux.Lock()
			var tmp int = counter
			tmp--
			counter = tmp
			mux.Unlock()
			wg.Done()
		}()
		fmt.Println("parallelMutex goroutine #'s: ", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("parallelMutex Counter: ", counter)
}

func parallelAtomic() {
	var counter int32 = 100
	const totit = 100
	var wg sync.WaitGroup

	wg.Add(totit)
	for i := 0; i < totit; i++ {
		go func() {
			atomic.AddInt32(&counter, -1)
			runtime.Gosched()
			wg.Done()
		}()
		fmt.Println("parallelAtomic gorutine #'s", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("parallelAtomic Counter", counter)
}
