package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Printf("Goroutines running %d\n", runtime.NumGoroutine())
	ctx := context.Background()
	ctx, cancelFn := context.WithCancel(ctx)
	defer cancelFn()

	for i := 0; i < 40; i++ {
		go func(n int) {
			fmt.Println("Launching goroutine: ", n)
			for {
				select {
				case <-ctx.Done():
					runtime.Goexit()
					// return
				default:
					fmt.Printf("goroutine %d is doing work\n", n)
					time.Sleep(100 * time.Millisecond)
				}
			}
		}(i)
	}

	time.Sleep(time.Millisecond)
	fmt.Printf("Goroutines running after one millisecond!!! %d\n", runtime.NumGoroutine())
	cancelFn()
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Goroutines running after cancelFn called!!! %d\n", runtime.NumGoroutine())
}
