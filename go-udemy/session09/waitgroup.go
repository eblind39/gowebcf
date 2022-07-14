package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("\t\tOS: ", runtime.GOOS)
	fmt.Println("\t\tARCH: ", runtime.GOARCH)
	fmt.Println("\t\tCPUs: ", runtime.NumCPU())
	fmt.Println("\t\tGORutines: ", runtime.NumGoroutine())

	wg.Add(2)

	go foo()

	result := make(chan int)
	go func() {
		result <- bar(10)
		wg.Done()
	}()
	fmt.Println("bar: ", <-result)

	wg.Wait()


	// Race condition -
	// To detect race conditions in code:
	// $  go build -race waitgroup.go
	// and exec file
	var contador int = 0
	var wgwg sync.WaitGroup
	var mux sync.Mutex
	const gs = 100
	wgwg.Add(gs)
	for i:=0; i<gs; i++ {
		go func() {
			mux.Lock()
			var v int = contador
			v++
			runtime.Gosched()
			contador = v
			mux.Unlock()
			wgwg.Done()
		}()
		fmt.Println("Número de Gorutinas ", runtime.NumGoroutine())
	}
	wgwg.Wait()
	fmt.Println("Cuenta: ", contador)
	// Race condition -

	// Race condition -
	// To detect race conditions in code:
	// $  go build -race waitgroup.go
	// and exec file
	var counter int64 = 0
	var wgwgwg sync.WaitGroup
	wgwgwg.Add(gs)
	for i:=0; i<gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("Counter: ", atomic.LoadInt64(&counter))
			wgwgwg.Done()
		}()
		fmt.Println("Número de Gorutinas ", runtime.NumGoroutine())
	}
	wgwgwg.Wait()
	fmt.Println("Cuenta: ", counter)
	// Race condition -
}

func foo() {
	for i:=0; i<10; i++ {
		fmt.Println("foo: ", i)
	}
	wg.Done()
}

func bar(number int) int {
	for i:=0; i<10000; i++ {
		number = int(math.Pow(float64(number), 2))
		number = int(math.Pow(float64(number), 0.5))
	}
	return number
}